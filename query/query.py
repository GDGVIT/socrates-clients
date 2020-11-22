import random
import os
import pickle

from scholarly import scholarly
from config import get_log_path

def first_unique_query(prevQueries, searchQuery):
    """Selects the first query which has not been served to user before"""
    pub = next(searchQuery)
    while True:
        if pub.bib['title'] not in prevQueries:
            break
        pub = next(searchQuery)
        
    return pub

def generate_query_string(topics):
    """Generates search string from a list of topics"""
    # Randomly pick half the topics from the list to make query
    n = int(len(topics)/2)
    indices = random.sample(range(0, len(topics)), n)

    queryTopics = [topics[i] for i in indices] 
    queryString = ''

    for topic in queryTopics:
        queryString += topic + ' '

    return queryString

def make_query(topics):
    """Queries Google scholar and returns the first new research paper found according to topics"""
    queryString = generate_query_string(topics)
    
    for topic in topics:
        queryString += topic + ' '
    
    logsDir = get_log_path()
    queryListPath = os.path.join(logsDir, 'query-list.pickle')

    # If any queries made previously, load that list, else create a new empty list 
    try:
        with open(queryListPath, 'rb') as f:
            prevQueries = pickle.load(f)
    except FileNotFoundError:
        prevQueries = []

    searchQuery = scholarly.search_pubs(queryString)
    pub = first_unique_query(prevQueries, searchQuery)

    prevQueries.append(pub.bib['title'])

    with open(queryListPath, 'wb+') as f:
        pickle.dump(prevQueries, f)

    return pub