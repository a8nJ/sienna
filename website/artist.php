<?php
declare(strict_types = 1);
$s_file = $_GET['f'];
$s_artist = $_GET['a'];
?>
<head>
   <link rel="icon" href="/sienna.png">
   <link rel="stylesheet" href="/sienna.css">
   <title><?= $s_artist ?> - Sienna</title>
</head>
<body>
   <header>
      <a href="..">Up</a>
<?php
echo <<<eof
<a href="/remote.php?f=$s_file&a=$s_artist">Remote</a>
eof;
?>
      <h1><?= $s_artist ?></h1>
   </header>
   <table>
<?php
$s_get = file_get_contents('../json/' . $s_file);
$o_get = json_decode($s_get);
foreach ($o_get->$s_artist as $s_album => $o_album) {
   $s_date = $o_album->{'@date'};
echo <<<eof
<tr>
   <td>$s_date</td>
   <td>
      <a href="/album.php?f=$s_file&a=$s_artist&r=$s_album">$s_album</a>
   </td>
</tr>
eof;
}
?>
   </table>
</body>
