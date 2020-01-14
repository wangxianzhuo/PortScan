package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
	"github.com/wangxianzhuo/PortScan/config"
	"github.com/wangxianzhuo/PortScan/output"
)

var (
	configFile           = flag.String("config", "./config.json", "配置文件路径")
	scanCron             = flag.String("cron", "@every 1m", "扫描周期的cron表达式")
	scanTimeout          = flag.Int("timeout", 1, "扫描超时设置，单位秒（s）")
	maxConcurrentScanNum = flag.Int("max-con-num", 5, "同时扫描端口最大个数")
)

func main() {
	flag.Parse()

	// 加载配置
	conf, err := config.Load(*configFile)
	if err != nil {
		log.Panic("加载配置失败：", err)
	}
	fmt.Println("###################################################################################")
	log.Println("配置加载\t\t\t完成")
	log.Printf("配置文件路径\t\t%s", *configFile)
	log.Printf("扫描周期的cron\t\t%s", *scanCron)
	log.Printf("扫描超时\t\t\t%ds", *scanTimeout)
	log.Printf("同时扫描个数\t\t%d个", *maxConcurrentScanNum)
	fmt.Println("###################################################################################")
	fmt.Println("##")
	fmt.Println("## 扫描列表")
	fmt.Println("##")
	for _, scanObj := range conf.ScanList {
		fmt.Printf("## " + scanObj.IP + ": [")
		for _, port := range scanObj.Ports {
			fmt.Printf(" %d ", port)
		}
		fmt.Printf("]\n")
	}
	fmt.Println("##")
	fmt.Println("###################################################################################")

	scanRoutineLimit := make(chan struct{}, *maxConcurrentScanNum)

	// 加载输出器实现
	outputMap := make(map[string]string)
	for _, outputObj := range conf.OutputList {
		outputMap[outputObj.Name] = outputObj.Ref
	}
	outputers := output.Load(outputMap)
	log.Println("输出器加载\t\t\t完成")

	// 创建cron任务
	c := cron.New()
	c.AddFunc(*scanCron, func() {
		for _, scanObj := range conf.ScanList {
			for _, port := range scanObj.Ports {
				scanRoutineLimit <- struct{}{}
				go func(ip string, port int) {
					// 扫描端口
					err := scanPort(ip, port)
					if err != nil {
						output.All(outputers, fmt.Sprint(err), fmt.Sprintf("%s:%d", ip, port))
					}
					<-scanRoutineLimit
				}(scanObj.IP, port)
			}
		}
	})

	c.Start()
	log.Println("定时任务\t\t\t启动")
	fmt.Println("###################################################################################")

	// 监听退出信号
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSTOP)
	go func() {
		sig := <-sigs
		fmt.Println()
		log.Printf("接收控制信号\t\t%s", sig)
		done <- true
	}()
	<-done
	log.Println("端口扫描程序\t\t退出")
}

func scanPort(ip string, port int) error {
	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}

	log.Printf("扫描开始\t\t\t扫描: %s", tcpAddr.String())
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), time.Duration(*scanTimeout)*time.Second)
	if err != nil {
		return errors.New(err.Error() + fmt.Sprintf(", 超时时间: %s", time.Duration(*scanTimeout)*time.Second))
	}
	defer conn.Close()
	log.Printf("扫描结束，端口已开启\t扫描: %s", tcpAddr.String())
	return nil
}
