package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron"
	"github.com/wangxianzhuo/PortScan/config"
	"github.com/wangxianzhuo/PortScan/output"
)

var (
	configFile = flag.String("config", "./config.json", "配置文件路径")
	scanCron   = flag.String("cron", "@every 1m", "扫描周期的cron表达式")
)

func main() {
	flag.Parse()

	// 加载配置
	conf, err := config.Load(*configFile)
	if err != nil {
		log.Panic("加载配置失败：", err)
	}
	log.Println("配置加载，完成")

	// 加载输出器实现
	outputMap := make(map[string]string)
	for _, outputObj := range conf.OutputList {
		outputMap[outputObj.Name] = outputObj.Ref
	}
	outputers := output.Load(outputMap)
	log.Println("输出器加载，完成")

	// 创建cron任务
	c := cron.New()
	c.AddFunc(*scanCron, func() {
		for _, scanObj := range conf.ScanList {
			for _, port := range scanObj.Port {
				err := scanPort(scanObj.IP, port)
				if err != nil {
					output.All(outputers, fmt.Sprint(err))
				}
			}
		}
	})

	c.Start()
	log.Println("定时任务，启动")

	// 监听退出信号
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSTOP)
	go func() {
		sig := <-sigs
		fmt.Println()
		log.Println(sig)
		done <- true
	}()
	<-done
	log.Println("端口扫描程序，退出")
}

func scanPort(ip, port string) error {
	return nil
}
