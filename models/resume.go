package models

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/twinj/uuid"
	"gopkg.in/gomail.v2"
)

type Powerby struct {
	Application   string            `json:"application"`
	Version       string            `json:"version"`
	Released      string            `json:"released"`
	Documentation string            `json:"documentation"`
	Licensed      string            `json:"licensed"`
	CreatedBy     map[string]string `json:"created_by"`
	Other         string            `json:"other"`
}

type Personal struct {
	Name     string `json:"name"`
	Job      string `json:"job"`
	Newjob   string `json:"newjob"`
	Birthday string `json:"birthday"`
	Website  string `json:"website"`
	Degree   string `json:"degree"`
	Addr     string `json:"addr"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	Age      uint8  `json:"age"`
	Img      string `json:"img"`
}

type Interest struct {
	Name string `json:"name"`
}

type Description struct {
	Detail string `json:"detail"`
}

type Experience struct {
	Organization   string        `json:"organization"`
	Job            string        `json:"job"`
	Sd             string        `json:"sd"`
	Ed             string        `json:"ed"`
	Detail         string        `json:"detail"`
	JobDescription []Description `json:"job_description"`
}

type Education struct {
	Sd         string `json:"sd"`
	Ed         string `json:"ed"`
	Degree     string `json:"degree"`
	Major      string `json:"major"`
	University string `json:"university"`
	Detail     string `json:"detail"`
}

type Skill struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

type Tool struct {
	Name string `json:"name"`
}

type About struct {
	Detail1 string `json:"detail1"`
	Detail2 string `json:"detail2"`
}

type Reference struct {
	Comment string `json:"comment"`
	Img     string `json:"img"`
	Name    string `json:"name"`
	Job     string `json:"job"`
}

type Portfolio struct {
	System string `json:"system"`
	Detail string `json:"detail"`
}

type Contact struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
	CreateAt string `json:"create_at"`
}

func (c *Contact) TableName() string {
	return "contacts"
}

func HealthCheck(host string) Powerby {
	var powerby = make(map[string]string)

	url := host + "/swagger/index.html"

	powerby["name"] = "Mr.Navamin Sawasdee"
	powerby["tel"] = "082-7893384"
	powerby["email"] = "navaminsawasdee@gmail.com"

	p := Powerby{}
	p.Application = "DEMO API V1"
	p.Version = "v1.0.0."
	p.Released = "2022-11-11"
	p.Documentation = url
	p.Licensed = "Copyright © 2022 By Navamin Sawasdee. All rights reserved."
	p.CreatedBy = powerby
	p.Other = "Developing a RESTful API with Go 1.17 and Gin Framework"

	return p
}

func SkillsInfo() []Skill {

	var skill = []Skill{
		{Name: "Golang", Class: "bx bxl-go-lang bx-lg"},
		{Name: "Gin framework", Class: "bx bxl-go-lang bx-lg"},
		{Name: "Laravel framework", Class: "bx bxl-php bx-lg"},
		{Name: "Bash", Class: "bx bx-terminal bx-md"},
		{Name: "Swagger", Class: "bx bx-book bx-md"},
		{Name: "MySQL", Class: "bx bxs-cylinder bx-md"},
		{Name: "Redis", Class: "bx bxs-cylinder bx-md"},
		{Name: "Nginx", Class: "bx bx-cog bx-md"},
		{Name: "Apache", Class: "bx bx-cog bx-md"},
		{Name: "Linux", Class: "bx bx-terminal bx-md"},
		{Name: "Cloud(Pass)", Class: "bx bx-terminal bx-md"},
		{Name: "Web Service", Class: "bx bx-cog bx-md"},
		{Name: "Web Application", Class: "bx bx-cog bx-md"},
		{Name: "System design", Class: "bx bx-cog bx-md"},
		{Name: "Software design", Class: "bx bx-cog bx-md"},
	}

	return skill
}

func ToolsInfo() []Tool {

	var tool = []Tool{
		{Name: "Visual Studio Code"},
		{Name: "Postman"},
		{Name: "REST Client"},
		{Name: "MySQL Workbench"},
		{Name: "TortoiseSVN"},
		{Name: "WinSCP"},
		{Name: "Git"},
		{Name: "PuTTY"},
		{Name: "Microsoft Teams"},
		{Name: "Slack"},
		{Name: "VMware Workstation"},
		{Name: "DigitalOcean(PaaS)"},
	}

	return tool
}

func PersonalInfo(host string) Personal {
	var personal Personal

	personal.Name = "Navamin Sawasdee"
	personal.Newjob = "Senior Software Engineer(Golang)"
	personal.Birthday = "7 June 1991"
	personal.Website = host
	personal.Tel = "082-789-3384"
	personal.Addr = "Krung Thep Maha Nakhon"
	personal.Age = 31
	personal.Degree = "Bachelor of Science(B.Sc.)"
	personal.Email = "navaminsawasdee@gmail.com"
	personal.Job = "Senior Software Engineer"
	personal.Img = "assets/img/IMG_9218.jpg"
	return personal
}

func ExperienceInfo() []Experience {
	var ex = []Experience{
		{
			Organization: "T.N. INCORPORATION LTD, Thailand",
			Job:          "Senior Software Engineer",
			Sd:           "2022",
			Ed:           "Present",
			Detail:       "T.N. was the first Thai company to offer IT consulting specialised in mission-critical, large-scale and complex system. TN is a market leader amongst large-scale banking solution providers with the most market share in Core Banking System implementation.",
			JobDescription: []Description{
				{Detail: "Develop Banking system with TN Golang framework"},
			},
		},
		{
			Organization: "Government Savings Bank(GSB), Thailand",
			Job:          "Software Developer",
			Sd:           "2017",
			Ed:           "2022",
			Detail:       "Government Savings Bank (GSB) provides savings, credit and other financial services to personal customer group, grassroots and government policy's customer group, and business and public sectors customer group.",
			JobDescription: []Description{
				{Detail: "Design and develop Recurring system, Event Management Smart Application, Aoon-i-rak, Children's Day, Salak Management with Laravel framework, Time Attendance API with Gin framework."},
				{Detail: "Ability to write technical documents."},
				{Detail: "Participate in support of Credit card system and Line Chatbot."},
				{Detail: "Contribute to the development of customer interfaces."},
			},
		},
		{
			Organization: "Triple V Broadcast Co., Ltd.(Thairath TV), Thailand",
			Job:          "Broadcast Software Specialist",
			Sd:           "2016",
			Ed:           "2017",
			Detail:       "Thairath TV is a Digital Terrestrial Television owned by the news publisher, Thairath, launched in April 2014 after they obtained a digital television broadcast license.",
			JobDescription: []Description{
				{Detail: "Implemented and maintained Octopus System and Mosart System (News Automation)."},
				{Detail: "Troubleshoot Octopus System, Mosart System and Video server."},
				{Detail: "Troubleshooting a variety of hardware, software and network issues."},
				{Detail: "Managed employee development and training program."},
			},
		},
		{
			Organization: "Triple V Broadcast Co., Ltd.(Thairath TV), Thailand",
			Job:          "News Generation Technology",
			Sd:           "2014",
			Ed:           "2016",
			Detail:       "Thairath TV is a Digital Terrestrial Television owned by the news publisher, Thairath, launched in April 2014 after they won a digital television broadcast license.",
			JobDescription: []Description{
				{Detail: "Implemented, maintained and troubleshoot Media Asset Management System (MAM) and Pebble Beach System."},
				{Detail: "Troubleshooting a variety of hardware, software and network issues."},
				{Detail: "Managed employee development and training program."},
			},
		},
	}
	return ex
}

func EducationInfo() Education {
	var ed = Education{
		Sd:         "2010",
		Ed:         "2014 ",
		Degree:     "Bachelor's Degree of  Science (B.Sc.)",
		Major:      "Computer Information System",
		University: "Burapha University, Thailand",
		Detail:     "Graduated from Burapha University in August 2014 with a B.Sc. in Computer Information System.Internships and develop Register System Electricity Generating Authority of Thailand at Chachoengsao.",
	}
	return ed
}

func AboutInfo() About {
	var intro = About{
		Detail1: "Hello, I'm Navamin Sawasdee you can call me Jane. I started working as News Generation Technology at Triple V Broadcast Co., Ltd.(Thairath TV) for about 3 years. I gained a veriety of technical knowledge including memorable experiena there. After that I move to work for Government Savings Bank(GSB) and T.N. INCORPORATION LTD. I like critical thinking and solving problems. I can cope with the project with my backend web development experience. However, Learning new things always my favorite, I started programming in PHP but I enjoy Golang the most.",
		Detail2: "I am interested in a Back-End positions mainly involves Go. I'm mostly interested in Microservices, Programming Language Implementations and Cloud Systems. I would like to work in a Krung Thep Maha Nakhon.",
	}
	return intro
}

func InterestInfo() []Interest {
	var interest = []Interest{
		{Name: "Microservices"},
		{Name: "Kubernetes"},
		{Name: "Line Chatbot"},
		{Name: "RPA"},
		{Name: "Cloud"},
		{Name: "Football"},
		{Name: "Dota2"},
		{Name: "Valorant"},
	}

	return interest
}

func CreateContact(name, email, subject, msg string) error {
	contact := Contact{}
	id := uuid.NewV4().String()

	contact.Id = id
	contact.Name = name
	contact.Email = email
	contact.Subject = subject
	contact.Message = msg
	contact.CreateAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := contact.SaveContact()
	if err != nil {
		log.Error(err)
		return err
	} else {
		_, err = contact.SendContact()
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func (c *Contact) SaveContact() (*Contact, error) {
	err := DB.Create(&c).Error
	if err != nil {
		log.Error(err)
		return &Contact{}, err
	}
	return c, nil
}

func (c *Contact) SendContact() (bool, error) {

	from := "navaminsawasdee@gmail.com"
	subject := "Thanks for your interest in me"
	msg := fmt.Sprintf("Dear %s<br>I will contact you back later.<br><br>I can be reached anytime via email at %s or my cell phone, 082-789-3384.<br><br>Best regards,<br>%s", c.Name, from, "Navamin Sawasdee")

	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", c.Email)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", msg)

	// Send the email
	if err := Mailer.DialAndSend(message); err != nil {
		log.Error("[Mailer] ", err)
		return false, err
	}
	return true, nil
}
