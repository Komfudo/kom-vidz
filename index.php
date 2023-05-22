<html>
<head>
<title>KOM-VIDS</title>
<style>
table, th, td {
  border:1px solid white;
}
body {
  width: 35em;
  margin: 0 auto;
  background-color: black;
  color: white;
  font-family: Arial, Helvetica, sans-serif;
  text-align: center;
}
</style>
<link rel="icon" type="image/x-icon" href="/assets/icon.ico">
</head>
<body>
<img src="/assets/banner.png">
<p>Stupid videos made by me and also web design is my passion.<br>Select a video to download (for free)</p>
<?php

// Assuming you have already established a database connection
$host = "localhost";
$username = "root";
$password = "";
$dbname = "komvidz";

// Create a connection
$connection = mysqli_connect($host, $username, $password, $dbname);

// Check connection
if (!$connection) {
    die("Connection failed: " . mysqli_connect_error());
}

// Fetch video data from the database
$query = "SELECT * FROM Videos";
$result = mysqli_query($connection, $query);

if (mysqli_num_rows($result) > 0) {
    // If videos are found in the database, generate the table
    echo "<table>
          <tr>
            <th>Title</th>
            <th>Date</th>
            <th>Video file</th>
          </tr>";

    while ($row = mysqli_fetch_assoc($result)) {
        echo "<tr>
                <td>".$row['Title']."</td>
                <td>".$row['DateTime']."</td>
                <td><a href='".$row['FilePath']."' download>Download</a></td>
              </tr>";
    }

    echo "</table>";
} else {
    // If no videos are found in the database, display a message
    echo "<h1>No videos found in the database</h1>";
}

// Close the database connection
mysqli_close($connection);
?>
<p><em>vidz.komfudo.eu.org not copyrighted 2023</em></p>
</body>
</html>
