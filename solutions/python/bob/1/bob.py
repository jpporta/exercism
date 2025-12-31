def response(hey_bob):
    clean = hey_bob.strip()
    if clean == "":
        return "Fine. Be that way!"
    question = clean[-1] == '?'
    yell = clean == clean.upper() and any(c.isalpha() for c in clean)
    if yell and question:
        return "Calm down, I know what I'm doing!"
    if question:
        return "Sure."
    if yell:
        return "Whoa, chill out!"
    return "Whatever."




