import yaml

def get_yaml(path):
    with open(path, 'r') as stream:
        return yaml.safe_load(stream)

def get_port():
    path = os.path.join('..','settings','config.yaml')
    config = get_yaml(path)
    return config['API_PORT']

def get_log_path():
    path = 'config.yaml'
    config = get_yaml(path)
    return config['QUERY_LOGS_DIR']
