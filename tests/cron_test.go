/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"log"
	"testing"
)

func TestCronStart(t *testing.T) {
	Test.Start("Cron")
}

func testAddTask() {
	log.Println("Add task ok")
}

// Execute every five seconds
// func TestCronAddTask(t *testing.T) {
//	task1 := "*/5 * * * * *"
//	task2 := "*/10 * * * * *"
//	system.Quartz.New().
//		Add(task1, testAddTask).
//		Add(task2, testAddTask).
//		Start()
//}

func TestCronEnd(t *testing.T) {
	Test.End("Cron")
}
