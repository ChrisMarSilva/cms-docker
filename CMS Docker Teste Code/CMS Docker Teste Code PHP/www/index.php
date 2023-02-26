
<?php

$fileName = $_POST['txtFilaName'];

?>

<form action="" method="post">
  <label>Arquivo:</label>
  <input type="text" id="txtFilaName" name="txtFilaName" value="<?php echo $fileName;?>"><br><br>
  <input type="submit" value="Carregar">
</form>
  
<?php

class LerArquivo {

  private $conteudo = '';
  
  function exibirConteudo() {
    return $this->conteudo;
  }

  function carregar($fileName){
    $this->conteudo = '';
    try{

      if(!file_exists($fileName)){
        $this->conteudo = 'Arquivo não localizado<br>';
        return;
      }

      $file = fopen($fileName, "r");
      while ($line = fgets($file)) {
        $this->conteudo .= $line . "<br />";
      }
      fclose($file);

    } catch(Exception $e){
      return $e;
    }
  }

}

$arqv = new LerArquivo();
$arqv->carregar($fileName);
echo $arqv->exibirConteudo();

?>