package queue

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	tasks2 "github.com/s4kibs4mi/snapify/tasks"
	"github.com/s4kibs4mi/snapify/worker"
)

func SendTakeScreenShotTask(ID string) error {
	sig := &tasks.Signature{
		Name: tasks2.TakeScreenShotTaskName,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: ID,
			},
		},
	}
	_, err := worker.MachineryServer().SendTask(sig)
	if err != nil {
		return err
	}
	return nil
}
