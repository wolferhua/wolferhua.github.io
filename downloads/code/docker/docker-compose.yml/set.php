<?php

$f = fopen("yml.yml",'r');

function getDir($dir ,$spec)
{
    for ($i=0; $i < $spec; $i++) { 
       $dir = dirname($dir);
    }
    return $dir;
}
function getSpec( $spec)
{
    $ss = [];
    for ($i=0; $i < $spec; $i++) { 
       $ss[]="├──";
    }
    return implode("",$ss);
}

$base=__DIR__.'/x';
echo $base,"\n";

$spec = 0;
while ($line=fgets($f)) {
    $ret = preg_match('/^( *)([a-z_]+)\s*#(.*)/i',$line,$matches);
    // print_r($matches);
    if(count($matches)>3){
        $li = [
            "name"=>$matches[2],
            "desc"=>$matches[3],
            "spec"=>strlen($matches[1])/4,
        ];
         
        // echo $base,"\n";
        $spec =getSpec( $li['spec']);
        echo $spec;
        echo " ",$li["name"],"\t\t#",$li['desc'],"\n";
        
        // $base = $base.'/'.$li['name'];
        //mkdir($base,0777,true);
        // file_put_contents("${base}/desc",$li['desc']);
        // echo $base,"\n" ; 

    } 
   
}


