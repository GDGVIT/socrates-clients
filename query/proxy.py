from proxyscrape import create_collector

def generate_proxy(protocol):
    """Retrieves a proxy according to protocol and returns its address"""
    if protocol not in ['https', 'http']:
        return None

    collector = create_collector(protocol + '-collector', protocol)
    collected = collector.get_proxy()
    proxy = protocol + '://' + collected.host + ':' + collected.port
    return proxy


def get_proxy():
    """Get HTTP and HTTPS proxy addresses"""
    httpProxy = generate_proxy('http')
    httpsProxy = generate_proxy('https')

    return httpProxy, httpsProxy
