package models

import (
	"github.com/astaxie/beego/orm"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Topic struct {
	Id              int64 `orm:"auto"`
	Uid             int64
	Title           string
	Content         string    `orm:"size(5000)"`
	Attachment      string    `orm:"null"`
	Created         time.Time `orm:"index;null"`
	Updated         time.Time `orm:"index;null"`
	Views           int64     `orm:"index;null"`
	Author          string    `orm:"null"`
	ReplyTime       time.Time `orm:"index;null"`
	ReplyCount      int64     `orm:"null"`
	ReplyLastUserId int64     `orm:"null"`
	Category        string    `orm:"null"`
	Labels          string    `orm:"null"`
}

func AddTopic(title, content, category, label, attachment string) error {

	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"
	o := orm.NewOrm()

	topic := &Topic{
		Title:      title,
		Content:    content,
		Category:   category,
		Labels:     label,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return err
}

func GetAllTopic(cate, label string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topic := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topic)
	} else {
		_, err = qs.All(&topic)
	}

	return topic, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	topic.Labels = strings.TrimRight(strings.Replace(strings.Replace(topic.Labels, "#", " ", -1), "$", "", -1), " ")
	return topic, err
}

func ModifyTopic(tid, title, content, category, label, attachment string) error {
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"
	tidNum, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		return err
	}

	var oldCate, oldAttach string
	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}

	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Content = content
		topic.Category = category
		topic.Labels = label
		topic.Attachment = attachment
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}

	//删除旧附件
	if len(attachment) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}

	//更新分类统计
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
			if err != nil {
				return err
			}
		}
	}
	if len(category) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", category).One(cate)
		if err == nil {
			cate.TopicCount++
			_, err = o.Update(cate)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	var oldCate string
	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}

	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
			if err != nil {
				return err
			}
		}
	}

	return err
}
