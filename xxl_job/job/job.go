package job

import (
	"context"
	"errors"
	"fmt"
	"github.com/feixiaobo/go-xxl-job-client/v2"
	"github.com/feixiaobo/go-xxl-job-client/v2/logger"
	"strconv"
)

// TestJob1 doc
func TestJob1(ctx context.Context) error {
	//GetParam 格式是key=value,key=value
	name, ok := xxl.GetParam(ctx, "name")
	age, ok := xxl.GetParam(ctx, "age")
	fmt.Println("name :",name,ok)
	ageint ,err :=strconv.Atoi(age)
	if err!=nil{
		fmt.Println("atoi",err)
		logger.Info(ctx, "atoi error accur: ",err)
		return  errors.New("atoi err is accured "+err.Error())
	}
	fmt.Println("age :",ageint,ok)
	fmt.Printf("%s : hello world\n", name)
	fmt.Printf("%#v", ctx)
	logger.Info(ctx, "success")
	return errors.New("!!!!!!!!!!!!!!!!!!!! ")
}

func TestJob2(ctx context.Context) error {
	logger.Info(ctx, "golang job run success >>>>>>>>>>>>>>")
	logger.Info(ctx, "the input param:")
	return nil
}
