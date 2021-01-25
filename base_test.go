package xcore

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/plugin/dbresolver"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	Bootstrap("dev")
}

func TestDb(t *testing.T) {

	type Test struct {
		Id   string
		Name string
	}
	//tablePlus init
	/*
			CREATE DATABASE `m` DEFAULT CHARACTER SET = `utf8mb4` DEFAULT COLLATE = `utf8mb4_general_ci`;
		use m;
		CREATE TABLE`test`(`id` int(11)NOT NULL AUTO_INCREMENT,`name` varchar(255)NOT NULL,PRIMARY KEY(`id`))ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
		INSERT INTO `m`.`test` (`id`, `name`) VALUES ('1', 'd-master');


			CREATE DATABASE `a` DEFAULT CHARACTER SET = `utf8mb4` DEFAULT COLLATE = `utf8mb4_general_ci`;
		use `a`;
		CREATE TABLE`test`(`id` int(11)NOT NULL AUTO_INCREMENT,`name` varchar(255)NOT NULL,PRIMARY KEY(`id`))ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
		INSERT INTO `a`.`test` (`id`, `name`) VALUES ('1', 'a-master');
	*/

	//m-r
	res_w := &Test{}
	_ = G_DB.Table("test").Find(&res_w)
	fmt.Println("res_w,id=", res_w.Id, "name=", res_w.Name)
	//m-w
	res_r := &Test{}
	_ = G_DB.Clauses(dbresolver.Write).Table("test").Find(&res_r)
	fmt.Println("res_r,id=", res_r.Id, "name=", res_r.Name)

	//a-r
	ac_r := &Test{}
	_ = G_DB.Clauses(dbresolver.Use("activity")).Table("test").Find(&ac_r)
	fmt.Println("ac_r,id=", ac_r.Id, "name=", ac_r.Name)
	//a-w
	ac_w := &Test{}
	_ = G_DB.Clauses(dbresolver.Use("activity"), dbresolver.Write).Table("test").Find(&ac_w)
	fmt.Println("ac_w,id=", ac_w.Id, "name=", ac_w.Name)

	defer CloseRes()
}

func TestLog(t *testing.T) {
	G_LOG.Out = os.Stdout
	G_LOG.Formatter = &logrus.JSONFormatter{}
	G_LOG.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
	//G_LOG.SetFormatter(&logrus.JSONFormatter{})
	//G_LOG.SetReportCaller(true)
	//G_LOG.SetOutput(os.Stdout)
	//G_LOG.SetLevel(logrus.WarnLevel)
	//G_LOG.WithField("a-test","whahah").Info("a test log")
}

func TestRedis(t *testing.T) {
	Bootstrap("dev")
	fmt.Println(G_REDIS)
	G_REDIS.Set("ket", 1234, time.Duration(10)*time.Minute)
	fmt.Println(G_REDIS.Get("ket"))

}
