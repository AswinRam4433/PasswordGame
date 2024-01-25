# Password Game README

Welcome to the Password Game repository! This project is a backend implementation of the [Neal's Password Game](https://neal.fun/password-game/) challenge. The game involves creating a password that satisfies multiple rules. The rules range from simple requirements like password length and the presence of specific characters to more complex conditions involving Roman numerals, moon phases, and dynamic data from external APIs.

## Table of Contents

- [Overview](#overview)
- [Rules](#rules)
- [Setup](#setup)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Postman Test Cases](#postman-test-cases)
- [Contributing](#contributing)
- [License](#license)

## Overview

This Go application provides a password validation service based on a set of predefined rules. The rules cover various aspects of password creation, including length, character types, dynamic conditions, and external data validation.

## Rules

1. **Length Check**: Password must have at least 5 characters.
2. **Number Required**: Password must contain at least one number.
3. **Uppercase Letter**: Password must contain at least one uppercase letter.
4. **Special Character**: Password must contain at least one special character.
5. **Digit Sum**: Digits in the password must add up to 25.
6. **Month Presence**: Password must contain the name of a month.
7. **Roman Numeral**: Password must contain a Roman numeral.
8. **Sponsor Check**: Password must contain one of the sponsors: Pepsi, Shell, Starbucks.
9. **Roman Numeral Multiplication**: Roman numerals in the password must multiply to 35.
10. **Wordle Of The Day**: Password must contain the Wordle of the day (dynamic data from an API https://github.com/sbplat/wordle).
11. **Moon Phase Emoji**: Password must contain an emoji representing the moon phase.

**The rules we check for here are a subset of the actual rules of Neal's game**

## Setup

1. Install [Go](https://golang.org/doc/install).
2. Clone this repository:

   ```bash
   git clone https://github.com/AswinRam4433/PasswordGame
   ```

3. Change to the project directory:


4. Run the application:

   ```bash
   go run main.go
   ```

## Usage

Once the application is running, it exposes an HTTP endpoint for password validation. You can make requests to this endpoint using the provided Postman test cases or by sending HTTP requests programmatically.

## Endpoints

- **POST /checkpassword**: Validate a password against a specific rule.
  - Request Body:

    ```json
    {
    "password": "12345A$55Ma",
    "ruleNumber": "6"
    }
    ```

  - Response:

    ```text
    Received user: {Password:12345A$55Ma Rule:6}
    Digits Must Add Up to 25
    ```

## Postman Test Cases

Import the provided Postman collections to quickly test different password scenarios against the defined rules.

## Contributing

Feel free to contribute to this project by opening issues, providing suggestions, or submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Video Explanation🎬
https://www.youtube.com/watch?v=Q_G7Kfut4K4
