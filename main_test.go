package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"testing"
	"uiTestsGo/pages"
)

const (
	path = "/opt/homebrew/bin/geckodriver"
	port = 8080
)

func startWebDriver(t *testing.T) selenium.WebDriver {
	service, err := selenium.NewGeckoDriverService(path, port)
	if err != nil {
		t.Fatalf("Error starting GeckoDriver: %v", err)
	}

	t.Cleanup(func() {
		service.Stop()
	})

	caps := selenium.Capabilities{"browserName": "firefox"}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d", port))
	if err != nil {
		service.Stop()
		t.Fatalf("Error creating new remote: %v", err)
	}

	t.Cleanup(func() {
		wd.Quit()
	})

	return wd
}

func TestFormPositive(t *testing.T) {
	wd := startWebDriver(t)
	defer wd.Quit()

	page := pages.NewFormPage(wd)
	if err := page.OpenPage(); err != nil {
		t.Errorf("Failed to open page: %s", err)
	}

	if err := page.EnterName("John Connor"); err != nil {
		t.Errorf("EnterName failed: %s", err)
	}

	if err := page.EnterPassword("12345"); err != nil {
		t.Errorf("EnterPassword failed: %s", err)
	}

	if err := page.SelectDrinks("drink2", "drink3"); err != nil {
		t.Errorf("SelectDrinks failed: %s", err)
	}

	if err := page.SelectColor("color3"); err != nil {
		t.Errorf("SelectColor failed: %s", err)
	}

	if err := page.SelectAutomation("Yes"); err != nil {
		t.Errorf("SelectAutomation failed: %s", err)
	}

	if err := page.EnterEmail("mail@mail.com"); err != nil {
		t.Errorf("EnterEmail failed: %s", err)
	}

	if err := page.EnterMessage("Test message"); err != nil {
		t.Errorf("EnterMessage failed: %s", err)
	}

	if err := page.Submit(); err != nil {
		t.Errorf("Submit failed: %s", err)
	}

	alert, err := wd.AlertText()
	if err != nil {
		t.Errorf("Failed to get alert text: %s", err)
	}
	if alert != "Message received!" {
		t.Errorf("Alert expected: 'Message received!', got: %s", alert)
	}

	wd.AcceptAlert()
}

func TestFormNegative(t *testing.T) {
	wd := startWebDriver(t)
	defer wd.Quit()

	page := pages.NewFormPage(wd)
	if err := page.OpenPage(); err != nil {
		t.Errorf("Failed to open page: %s", err)
	}

	if err := page.EnterName("Frodo Baggins"); err != nil {
		t.Errorf("EnterName failed: %s", err)
	}

	if err := page.EnterPassword("12345"); err != nil {
		t.Errorf("EnterPassword failed: %s", err)
	}

	if err := page.SelectDrinks("drink2", "drink3"); err != nil {
		t.Errorf("SelectDrinks failed: %s", err)
	}

	if err := page.SelectColor("color3"); err != nil {
		t.Errorf("SelectColor failed: %s", err)
	}

	if err := page.SelectAutomation("Yes"); err != nil {
		t.Errorf("SelectAutomation failed: %s", err)
	}

	if err := page.EnterEmail("invalid_email"); err != nil {
		t.Errorf("EnterEmail failed: %s", err)
	}

	if err := page.EnterMessage("Test message"); err != nil {
		t.Errorf("EnterMessage failed: %s", err)
	}

	if err := page.Submit(); err != nil {
		t.Errorf("Submit failed: %s", err)
	}

	alert, err := wd.AlertText()
	if err != nil {
		t.Errorf("Failed to get alert: %s", err)
	}

	if alert != "Message received!" {
		t.Fatalf("Expected 'Message received!', got '%s'", alert)
	}
	wd.AcceptAlert()
}
