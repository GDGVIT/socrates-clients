from notifypy import Notify

def send_notification(pub):
    """Sends notification to OS level notification server based on selected publication"""
    message = pub.bib['title'] + '\n'
    message += 'Read at -\n'
    message += pub.bib['url']

    notification = Notify()
    notification.title = 'Your research paper pick for the day'
    notification.message = message

    notification.send()