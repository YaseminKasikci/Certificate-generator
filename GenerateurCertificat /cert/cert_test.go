package cert

import "testing"

// tdd = test driving development. Following the test which will start by failing

func TestValidCertData(t *testing.T) {
c, err := New("Golang", "Bob", "2018-05-31")
if err != nil {
	t.Errorf("Cert data should be valid. err=%v", err)
}
if c == nil {
	t.Errorf("Cert should be valid reference. got=nil")
}

if c.Course != "GOLANG COURSE" {
	t.Errorf("Course name is not valid. expected='GOLANG COURSE', got=%v", c.Course)
}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "azertyuiovjbedivhbdvjnkpsqjcpozeivliqebvqbvmlqjdnmiernqovdfdcufugvbiugvydrexfgh"
	_, err := New(course, "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on too long course name(course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "",  "2018-05-31")
	if err == nil {
		t.Error("Error should be returned on empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	name :="azertyuiovjbedivhbdvjnkpsqjcpozeivliqebvqbvmlqjdnmiernqovdfdcufugvbiugvydrexfgh"
	_, err := New("GOLANG", name,  "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on too long course name(name=%s)", name)
	}
}

func TestValidDate(t *testing.T) {
	date := "02-01-2006"
	_, err := New("Golang", "Bob", date)
	if err == nil {
		t.Errorf("Error should be returned on invalid date(%s)", date)
	}
}