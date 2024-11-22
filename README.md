# Parallel Scraper - Schneider Electric 2024
<div style="display: flex; justify-content: center; align-items: center;">
  
  <img src="https://github.com/user-attachments/assets/ae355d4a-3191-4913-87f6-bfd3840ca0b8" alt="selo"> 
  <img src="https://github.com/user-attachments/assets/6201a1f8-3da0-49ba-aa09-f6144f8b3a3d" alt="go">
  <img src="https://github.com/user-attachments/assets/8faaa599-23df-4a2e-bc55-284735f1e6b9" alt="reddit">
  
</div>


This application is a simple **scraper** for Reddit, developed in **Go**. It allows users to get Reddit posts related to a specific topic and extracts the comments associated with those posts. The application performs sentiment analysis (positive or negative) on the comments using **NLP** (Natural Language Processing). It uses **parallel scraping** for faster data collection and operates through a **console-based interface**.

## Technologies

- **Go (Golang)**: The programming language used to develop the application.
- **Colly**: A popular Go library for parallel web scraping.
- **Sentiment**: A Go library for Natural Language Processing, used for sentiment analysis (positive/negative) of the comments.
- **Console-based UI**: An interactive console application without the need for a graphical user interface.

## How the Application Works

1. **Reddit Search**: The user enters a topic (e.g., "Trump win"), and the application searches Reddit for posts containing that topic.
2. **Scraping Posts and Comments**: Using Colly, the application first scrapes Reddit search results for posts related to the topic, then visits each post to collect comments. 
3. **Sentiment Analysis**: The comments collected are processed using NLP to determine whether they are positive or negative.
4. **Displaying Results**: The results of sentiment analysis are displayed directly in the console.


