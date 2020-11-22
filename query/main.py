from scholarly import ProxyGenerator
import requests
import logging

from notify import send_notification
from proxy import get_proxy
from config import get_log_path, get_port

def main():
    
    # Register a proxy for scholar queries
    # Avoids google blocking user's IP
    httpProxy, httpsProxy = get_proxy()
    proxyGen = ProxyGenerator()
    proxyGen.SingleProxy(http=httpProxy, https=httpProxy)
    scholarly.use_proxy(proxyGen)

    # Setup logging
    logging.basicConfig(filename=get_log_path(), encoding='utf-8', level=logging.DEBUG)

    # Get API port address
    PORT = str(get_port())

    # Get current settings from API
    res = requests.get('http://localhost:' + PORT + '/view').json()
    topics = res['topics']

    # Abort if topics list empty
    if len(topics) == 0:
        logging.info('Topics list empty, aborting query')
        return

    # Get publication details
    pub = make_query(topics)

    # Send notification
    send_notification(pub)

if __name__ == '__main__':
    main()