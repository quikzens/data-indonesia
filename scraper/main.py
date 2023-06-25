from selenium.webdriver import Chrome
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import time

from province import scrape_provinces

start = time.time()
scrape_provinces()
end = time.time()
print(end - start)