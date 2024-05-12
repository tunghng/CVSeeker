from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
import time
import random   

class ScrapinCrawler:
    def __init__(self) -> None:
        self.s = Service(r"chromedriver.exe")
        self.driver = webdriver.Chrome(service=self.s)
        self.driver.maximize_window()
        pass
    def sendData(self, LinkedInURL):
        login = open('login.txt') 
        line = login.readlines()
        api_key = line[0]
        self.driver.get("https://docs.scrapin.io/endpoint/ExtractPersonDataProfile")
        time.sleep(random.randint(7,10))
        input_api_key = self.driver.find_element(by=By.XPATH, value="//input[@placeholder='Enter apikey']")
        input_api_key.send_keys(api_key)
        
        input_url = self.driver.find_element(by=By.XPATH, value="//input[@placeholder='Enter linkedinUrl']")
        input_url.send_keys(LinkedInURL)
        time.sleep(5)
        send_button = self.driver.find_element(by=By.XPATH, value="//button[@class='flex items-center justify-center w-16 h-9 text-white font-medium rounded-lg mouse-pointer disabled:opacity-70 bg-[#2AB673]']")
        send_button.click()
        time.sleep(10)
        result = self.driver.find_element(by=By.XPATH, value="//span[@class='language-json']")
        print(result.text)
        self.driver.quit()
        pass
    def getData(self):
        result = self.driver.find_element(by=By.XPATH, value="//span[@class='language-json']")
        self.driver.quit()
        return result.text
