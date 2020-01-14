CREATE TABLE public.tcp_port_errors
(
    id serial NOT NULL,
    address character varying(50) NOT NULL,
    msg text NOT NULL,
    created_time timestamp without time zone NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE public.tcp_port_errors
    OWNER to dev;
COMMENT ON TABLE public.tcp_port_errors
    IS 'tcp端口异常表';

COMMENT ON COLUMN public.tcp_port_errors.address
    IS '异常地址';

COMMENT ON COLUMN public.tcp_port_errors.msg
    IS '异常信息';

COMMENT ON COLUMN public.tcp_port_errors.created_time
    IS '发生时间';