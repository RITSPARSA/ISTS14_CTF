<html>
<h1>Admin login</h1>

<?php
function connect_to_db(){
	return new mysqli("localhost", "al-dba", "ThisIsMuhDBpassword!", "AdminLogin");
}

function login($conn, $username, $ptpass){
    $password = hash("sha512", $username . $ptpass);
	$statement = $conn->prepare("SELECT * FROM users WHERE username = ? AND password = ?");
	$statement->bind_param("ss", $username, $password);
	$statement->execute();
	$query = $statement->get_result();
    $error = $statement->error;

	if($error){
        $result = $error;
        return $result;
    }
    if ($query->num_rows > 0) {
        $result = $query->fetch_array(MYSQLI_ASSOC);
        //$result = $query;
        $query->free();
    }
    else
        $result = False;

	return $result;
}

if(isset($_POST['username']) && isset($_POST['password']) && $_POST['username'] != "" && $_POST['password'] != "")
{
	$conn = connect_to_db();
	$result = login($conn, $_POST['username'], $_POST['password']);
	if($result)
	{
		$admin = True;
		//read in base64 cipher text and print to screen
		echo "<p>Remember Bob, you have to base64 decode this before you can decrypt it with your RSA private key.</p>\n";
		echo "<p>";
		echo file_get_contents("../cipher.b64");
		echo "</p>";
	}
	else
	{
		$admin = False;
		echo "<p style=\"color:red\">Access denied.</p>";
	}
}

if(!$admin)
{
	echo '<form action="index.php" method="POST">
	<div style="margin-bottom: 10px;">
		<span>Username</span>
		<input id="user-field" name="username" type="text">
	</div>
	<div>
		<span>Password</span>
		<input id="pass-field" name="password" type="password">
	</div>
	<input type="submit" value="Login">
</form>';
}
?>

</html>
