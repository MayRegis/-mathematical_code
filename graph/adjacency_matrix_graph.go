package main

import (
  "fmt"
)
const(
  //tam é a quantidade de vértices do grafo
  tam = 4
)
//=======================================================================
func valPrimFase() ([tam][tam]int){
  fmt.Print("\n________________________________________________________________________________________________________________")
  fmt.Print("\nNOTA: ")
  fmt.Print("\n1 - Mude o valor da 'const' no código (O valor corresponde a quantidade total de vértices no gráfo).")
  fmt.Print("\n2 - Monte sua matriz adjacência no papel e insira os valores aqui.")
  fmt.Print("\n3 - Enumere de '0' a 'n', onde 'n' é o vértice final e 0 o inicial.")
  fmt.Println("\nEx.:")

  //exemplo matriz
  var ex[4][4] int
  var adj[tam][tam] int
  var value int
  ex[0][0]=0;	ex[0][1]=5;	ex[0][2]=7;	ex[0][3]=4;	
	ex[1][0]=0;	ex[1][1]=0;	ex[1][2]=1;	ex[1][3]=0;	
	ex[2][0]=0;	ex[2][1]=0;	ex[2][2]=0;	ex[2][3]=2;	
	ex[3][0]=0;	ex[3][1]=0;	ex[3][2]=0;	ex[3][3]=0;	

	fmt.Printf("     Matriz Adjacência\n    ___________________\n   |___|_0_|_1_|_2_|_3_|");
	//Exemplo de tabela adjacencia
	for i := 0; i < 4; i++ {
		fmt.Printf("\n   |_%d_|",i);
		for j := 0; j < 4; j++ {
			fmt.Printf("_%d_|", ex[i][j]);
		}
	}

  fmt.Print("\n________________________________________________________________________________________________________________")
  //entrada dos valores da matriz
  fmt.Print("\n\n")
  for lin:=0;lin<tam;lin++{
    for col:=0;col<tam;col++{
      fmt.Print("adj[",lin,"][",col,"]=")
      fmt.Scanf("%d",&value)
      adj[lin][col]=value
    }
  } 
  
  fmt.Println("\nIniciando análise da Matriz de Adjacência...")

  return adj
}
//=======================================================================
func proxPosicao(col int, adj[tam][tam] int) (int, int, int) {
  var lin int
  var l int
  var c int
  var maior int

  //calculo da proxima posicao na matriz 
  for lin=0;lin<tam;lin++{
    if adj[lin][col] > 0{
      if adj[lin][col] > maior{
        maior=adj[lin][col]
        l=lin
        c=col
      }
    }
  }

  return l, c, maior
}
//=======================================================================
func PrimFase() ([tam][tam] int, int, [tam]int){
  var adj[tam][tam] int
  var col, lin, end int
  var l[tam] int
  var c[tam] int
  var maior[tam] int

  //Calculo da 1ª fase
  adj = valPrimFase()
  fmt.Print("\n________1ª fase_______\n")

  //seleciona a linha, coluna e o maior valor da ultima coluna
  for lin=0;lin<tam;lin++{
    for col=tam-1;col<tam;col++{
      if adj[lin][col] > 0 {
        if adj[lin][tam-1] > maior[0]{
          maior[0]=adj[lin][tam-1]
          l[0]=lin
          c[0]=col
        }
      }
    }
  }

  //print da matriz
  for lin:=0;lin<tam;lin++{
    fmt.Print("\n")
    for col:=0;col<tam;col++{
      fmt.Print(adj[lin][col]," ")
    }
  }

  //calcula a proxima posicao na matriz
  for i:=1;i<tam;i++{
    l[i],c[i],maior[i] = proxPosicao(l[i-1], adj)
  }
  
  //calcula o valor final
  adj, end = calcFinal(l,c,maior,adj)
  
  //print das posicoes usadas para calculo
  for i:=0;i<tam;i++{
    fmt.Printf("\nadj[%d][%d] = %d", l[i], c[i], maior[i])
  }
  
  return adj, end, c
}
//=======================================================================
func ProxFase(adj[tam][tam] int) ([tam][tam] int, int){
  var col, lin, end int
  var l[tam] int
  var c[tam] int
  var maior[tam] int

  //seleciona a linha, coluna e o maior valor da ultima coluna
  for lin=0;lin<tam;lin++{
    for col=tam-1;col<tam;col++{
      if adj[lin][col] > 0 {
        if adj[lin][tam-1] > maior[0]{
          maior[0]=adj[lin][tam-1]
          l[0]=lin
          c[0]=col
        }
      }
    }
  }
  
  //print da matriz
  for lin:=0;lin<tam;lin++{
    fmt.Print("\n")
    for col:=0;col<tam;col++{
      fmt.Print(adj[lin][col]," ")
    }
  }

  //calcula a proxima posicao na matriz
  for i:=1;i<tam;i++{
    l[i],c[i],maior[i] = proxPosicao(l[i-1], adj)
  }
  
  //calcula o valor final
  adj, end = calcFinal(l,c,maior,adj)
  
  //print das posicoes usadas para calculo
  for i:=0;i<tam;i++{
    fmt.Printf("\nadj[%d][%d] = %d", l[i], c[i], maior[i])
  }
  
  return adj, end
}
//=======================================================================
func calcFinal(l[tam] int, c[tam] int, maior[tam] int,  adj[tam][tam] int) ([tam][tam]int, int){
  var total[1000] int
  menor:=1000

  if maior[0] < menor{
      menor = maior[0]
  }

  for i:=0;i<tam;i++{
    if maior[0] > menor{
      total[i] = maior[0] - menor
    }else{
      total[i] = menor - maior[0]
    }  
    adj[l[i]][c[i]] = total[i]
  }  

  return adj, menor
}
//=======================================================================
func main() {
  var adj[tam][tam] int
  var mat[tam][tam] int
  var end[tam] int
  var value int

  //matriz e valor final da 1ª fase
  adj, end[0], _ = PrimFase()
  
  //matriz e valor final da 2ª fase
  fmt.Print("\n________2ª fase_______\n")
  mat, end[1] = ProxFase(adj)
  
  //matriz e valor final das demais fase
  for i:=2;i<tam;i++{
    if mat[i-1] != mat[i]{
      fmt.Print("\n________",i+1,"ª fase_______\n")
      mat, end[i] = ProxFase(mat)
    }else{
      fmt.Print("\n\nError: ",tam-1,"ª fase não adicionada!\n\nTodas arestas ja foram saturadas!!!") 
    }
  }

  //somatorio dos valores finais  
  for _, num := range end {
      value += num
  }

  fmt.Print("\n\n",end,"\nFluxo máximo: ", value," ")
}