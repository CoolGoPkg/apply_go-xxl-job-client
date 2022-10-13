package xxl_job

import (
	"CoolGoPkg/apply_xxl_job/conf"
	"CoolGoPkg/apply_xxl_job/xxl_job/job"
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/feixiaobo/go-xxl-job-client/v2"
	"github.com/feixiaobo/go-xxl-job-client/v2/option"
)

func StartXXLJobClient(config *conf.DemoConfig) {
	fmt.Println(config.XXLJobConf.AdminAddress)
	fmt.Println(config.XXLJobConf.AppName)
	fmt.Println(config.XXLJobConf.AppName)
	client := xxl.NewXxlClient(
		//option.WithEnableHttp(true), //xxl_job v2.2之后的版本
		//option.WithClientPort(8083),
		//option.WithAccessToken("default_token"),
		option.WithAppName(config.XXLJobConf.AppName),
		option.WithClientPort(config.XXLJobConf.ClientPort),
		option.WithAdminAddress(config.XXLJobConf.AdminAddress),
		//option.WithAdminAddress("http://127.0.0.1:8088/xxl-job-admin"),
	)
	client.SetLogger(&logrus.Entry{

		Logger: logrus.New(),
		Level:  logrus.InfoLevel,
	})
	// 定时任务
	// 退市股票最后一个交易日发送邮件提醒
	client.RegisterJob("TestJob1", job.TestJob1)
	client.RegisterJob("TestJob2", job.TestJob2)
	client.Run()
}
