<?php
declare(strict_types = 1);
?>
<link rel="stylesheet" href="/sienna.css">
<header>
   <div>
      <a href="..">Up</a>
   </div>
</header>
<main>
<?php
$s_file = $_GET['f'];
$s_get = file_get_contents('../json/' . $s_file);
$o_get = json_decode($s_get);

foreach ($o_get as $s_artist => $o_artist) {
   echo <<<eof
<div>
   <a href="/artist.php?f=$s_file&a=$s_artist">$s_artist</a>
</div>

eof;
}
?>
</main>