# UI Tests for Form Fields 
This project contains automated UI tests for the [Practice Automation Form Fields] (https://practice-automation.com/form-fields/) page, using **Go** and **Selenium WebDriver** with **Firefox (GeckoDriver)**.

**Note:** Analogous UI tests are also available in **Python** in a https://github.com/chinaeffrec/ui-tests-python.

## Test Scenarios

### 1. Positive Test: TestFormPositive
- Opens the form page. 
- Fills in:
     - Name 
     - Password 
     - Drinks (checkboxes)
     - Color (radio button)
     - Automation selection (dropdown)
     - Email 
     - Message 
- Submits the form. 
- Verifies that the alert displays: ```Message received!```

### 2. Negative Test: TestFormNegative
- Opens the form page. 
- Fills in the same fields, but with an invalid email. 
- Submits the form. 
- Verifies that the alert still displays: ```Message received!```

## Prerequisites
- Go installed (v1.20+ recommended)
- Firefox installed 
- GeckoDriver installed and accessible in PATH

## Running Tests
1. Navigate to the project root: ```cd uiTests```
2. Run all tests: ```go run run_tests.go```
