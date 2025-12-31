def is_armstrong_number(number):
    sum = 0
    for letter in str(number):
        sum += int(letter) ** len(str(number))
    return number == sum



