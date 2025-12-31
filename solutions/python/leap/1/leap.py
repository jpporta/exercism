def isDiv(n,d):
    return n % d == 0
def leap_year(year):
    return (year % 4 == 0) and not ((year % 100 == 0) ^ (year % 400 == 0)) 
