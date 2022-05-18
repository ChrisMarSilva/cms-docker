<?php
   
# 
# implementa funcao de calculo de pre�os e prazos 
# para servi�os dos correios
#
if( !function_exists( 'calculaFrete' ))
{
   
}

function calculaFrete(
      $cod_servico, /* codigo do servico desejado */
      $cep_origem,  /* cep de origem, apenas numeros */
      $cep_destino, /* cep de destino, apenas numeros */
      $peso,        /* valor dado em Kg incluindo a embalagem. 0.1, 0.3, 1, 2 ,3 , 4 */
      $altura,      /* altura do produto em cm incluindo a embalagem */
      $largura,     /* altura do produto em cm incluindo a embalagem */
      $comprimento, /* comprimento do produto incluindo embalagem em cm */
      $valor_declarado='0' /* indicar 0 caso nao queira o valor declarado */
   ){

      $cod_servico = strtoupper( $cod_servico );
      if( $cod_servico == 'SEDEX10' ) $cod_servico = 40215 ; 
      if( $cod_servico == 'SEDEXACOBRAR' ) $cod_servico = 40045 ; 
      if( $cod_servico == 'SEDEX' ) $cod_servico = 40010 ; 
      if( $cod_servico == 'PAC' ) $cod_servico = 41106 ;

      # ###########################################
      # C�digo dos Principais Servi�os dos Correios
      # 41106 PAC sem contrato
      # 40010 SEDEX sem contrato
      # 40045 SEDEX a Cobrar, sem contrato
      # 40215 SEDEX 10, sem contrato
      # ###########################################

      $correios = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?nCdEmpresa=&sDsSenha=&sCepOrigem=".$cep_origem."&sCepDestino=".$cep_destino."&nVlPeso=".$peso."&nCdFormato=1&nVlComprimento=".$comprimento."&nVlAltura=".$altura."&nVlLargura=".$largura."&sCdMaoPropria=n&nVlValorDeclarado=".$valor_declarado."&sCdAvisoRecebimento=n&nCdServico=".$cod_servico."&nVlDiametro=0&StrRetorno=xml";

      $xml = simplexml_load_file($correios);

      $_arr_ = array();
      if($xml->cServico->Erro == '0'):
         $_arr_['codigo'] = $xml -> cServico -> Codigo ;
         $_arr_['valor'] = $xml -> cServico -> Valor ;
         $_arr_['prazo'] = $xml -> cServico -> PrazoEntrega .' Dias' ;
         // return $xml->cServico->Valor;
         return $_arr_ ; 
      else:
         return false;
      endif;
   }

$origem = $_GET['origem'];
$destino = $_GET['destino'];
$peso = $_GET['peso'];
$altura = $_GET['altura'];
$largura = $_GET['largura'];
$comprimento = $_GET['comprimento'];
$servico = $_GET['servico'];

$_resultado = calculaFrete(  $servico,  $origem,  $destino,  $peso,  $altura, $largura, $comprimento, 0 );

echo "ORIGEM: ".$origem."<br>";
echo "DESTINO: ".$destino."<br>";
echo "CODIGO: ".$_resultado['codigo']."<br>";
echo "VALOR: ".$_resultado['valor']."<br>";
echo "PRAZO: ".$_resultado['prazo']."<br>";
echo "<br><br>". var_dump($_resultado)."<br><br>";

$_resultado2 = calculaFrete(  '41106',  '11680000',  '82220000',  '1',  '15', '22', '32', 0 );

echo "CODIGO: ".$_resultado['codigo']."<br>";
echo "VALOR: ".$_resultado2['valor']."<br>";
echo "PRAZO: ".$_resultado2['prazo']."<br><br>";
echo "<br><br>". var_dump($_resultado2)."<br><br>";

?>