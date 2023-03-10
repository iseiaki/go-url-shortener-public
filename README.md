<h1>Golang URL Shortener</h1>
<p>This is a simple web application that allows users to shorten the URL of a website. It was built using Golang as the back-end language and showcases my technical skills in back-end development, including:</p>
<ul>
  <li>Proficiency in Golang and SQL</li>
  <li>Working with servers</li>
  <li>Effective use of Golang frameworks - Gin, Gorm</li>
  <li>Knowledge of MySQL</li>
</ul>
<p>In addition to my back-end skills, this project also demonstrates my minimalist approach to design and user experience. By bringing together various elements of web development, I have created a fully functional website that delivers a seamless user experience.</p>
<p>I am dedicated to staying up-to-date with the latest web technologies and am committed to delivering high-quality results. Whether it's through implementing new features or fixing complex bugs, I am always striving to improve my skills and bring my projects to the next level.</p>
<h2>Technologies Used</h2>
<ul>
  <li>Golang</li>
  <li>SQL</li>
  <li>Gin framework</li>
  <li>Gorm framework</li>
  <li>MySQL</li>
  <li>HTML, CSS, JavaScript</li>
  <li>jQuery</li>
  <li>Bootstrap</li>
</ul>
<h2>Features</h2>
<ul>
  <li>URL validation to ensure correct input</li>
  <li>Clean and minimalist design for optimal user experience</li>
  <li>Functional URL changing service</li>
</ul>
<h2>API and how to use it</h2>
<ul>
<li>Send POST request to {site_domain}/api</li>
<li>Include the original url in the request body as JSON</li>
<li>Get the new URL back </li>
<li>Example in Python</li><br>
<pre><code>import requests
  url = "https://0r.wtf/api"
  data = {
      "url": "https://example.com"
  }
  response = requests.post(url, json=data)
  if response.status_code == 200:
      result = response.json()
      print(result["Newurl"])
  else:
      print("Request failed with status code", response.status_code)
      //Output: https://0r.wtf/{{ID}}</code></pre>
   
<h2>Deploy</h2>
Deploy can be seen at: <a href="https://0r.wtf" target="_blank">https://0r.wtf</a>

<h2>Getting Started</h2>
<ul>
  <li>Clone the repository</li>
  <li>Run the following command in the terminal to install dependencies:<br>
    <pre><code>go get -u github.com/gin-gonic/gin github.com/jinzhu/gorm github.com/go-sql-driver/mysql</code></pre>
  </li>
  <li>Create a new database and update the database details in the main.go file</li>
  <li>Run the following command in the terminal to start the application:<br>
    <pre><code>go run main.go</code></pre>
  </li>
  <li>Access the application in your browser at http://localhost:8080/</li>
</ul>
<h2>Contact</h2>
For any questions or feedback, feel free to reach out at mirokurtanidze16@gmail.com.
