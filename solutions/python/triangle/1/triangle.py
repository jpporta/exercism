def equilateral(sides):
    sides.sort(reverse=True)
    return len(set(sides)) <= 1 and sides[0] < sum(sides[1:])


def isosceles(sides):
    sides.sort(reverse=True)
    return len(set(sides)) <= 2 and sides[0] < sum(sides[1:])


def scalene(sides):
    sides.sort(reverse=True)
    return len(set(sides)) == 3 and sides[0] < sum(sides[1:])
