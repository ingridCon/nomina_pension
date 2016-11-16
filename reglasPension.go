package main

import (
  "fmt"
  ."github.com/mndrix/golog"
)

func main()  {
  m:= NewMachine().Consult(`
      tipo_pension(cc,pj).
      tipo_pension(cc,pi).
      tipo_pensionado(cc,pensionado_docente).
      tipo_pensionado(cc,pensionado_administrativo).
      tipo_pensionado(cc,pensionado_oficial).
      tipo_pensionado(cc,sustituto).

      conyugue(17078900,3).
      hijo(17078900,no).
      edad_person(1018,10).
      condicion(1018,estudiante).

      padre(17078900,5).
      madre(17078900,6).

      hermano(X,7).
      dependencia_economica(7,total).

      categoria_sustituto(X,C):-conyugue(X,Y),edad_person(Y,E),E@>=30,C = conyugue__beneficiario_vitalicio.
      categoria_sustituto(X,C):-conyugue(X,Y),edad_person(Y,E),E@<30,hijo(X,S),S==no,C = conyugue_beneficiario_temporal.
      categoria_sustituto(X,C):-conyugue(X,Y),edad_person(Y,E),E@<30,hijo(X,S),S==si,C = conyugue_beneficiario_vitalicio.
      categoria_sustituto(X,C):-hijo(X,H),edad_person(H,E),condicion(H,F),F==estudiante,E@>=18,E@=<25,C = hijo.
      categoria_sustituto(X,C):-hijo(X,H),edad_person(H,E),condicion(H,F),F==estudiante,E@<18,C = tutor_del_hijo.
      categoria_sustituto(X,C):-padre(X,P),hijo(X,V),conyugue(X,V),V==no, C = padre.
      categoria_sustituto(X,C):-madre(X,P),hijo(X,V),conyugue(X,V),V==no, C = madre.
      categoria_sustituto(X,C):-hermano(X,R),dependencia_economica(R,T),T==total, C = hermano_invalidez.

      valor_mesada(17078900,15896409).
      salario_minimo_legal(689400,2016).


      por_diez(D):-salario_minimo_legal(J,P), D is 10*J.
      por_veinte(V):-salario_minimo_legal(J,P), V is 20*J.

      aporte_fondoSoli(X,W):-valor_mesada(X,M),salario_minimo_legal(J,P),por_diez(D),por_veinte(V),M @>=D, M @=<V,  W is M*0.01.
      aporte_fondoSoli(X,W):-valor_mesada(X,M),salario_minimo_legal(J,P),por_veinte(V),M @>V, W is M*0.02.

      residencia(17078900,nacional).
      aporte_salud(X,C):-valor_mesada(X,Y),residencia(X,R),R == nacional, C is Y*0.12.
      aporte_salud(X,C):-valor_mesada(X,Y),residencia(X,R),R == extranjera, C is Y*0.125.

      valor_subsido_familiar(65900,2016).
      salario_minimo_conven(1752652,2016).
      tipo_pensionado(17078900,pensionado_administrativo).
      numero_beneficiarios(17078900,2).
      tipo_valor(tipo_A).

      subsidio_familiar(X,Y):-tipo_pensionado(X,P),valor_subsido_familiar(J,V),P==pensionado_administrativo, Y is J.
      subsidio_familiar(X,Y):-tipo_pensionado(X,P),salario_minimo_conven(J,V),P==pensionado_oficial, Y is J/30*2.55.
      pago_subfamiliar(X,F):-valor_mesada(X,M),salario_minimo_conven(J,P),subsidio_familiar(X,Y),numero_beneficiarios(X,H),M@<5*J, F is Y*H.

      pago_subsidio_libros(X,S):-tipo_pensionado(X,P),tipo_valor(T),P==pensionado_administrativo,T==tipo_A,salario_minimo_conven(J,U),S is J.
      pago_subsidio_libros(X,S):-tipo_pensionado(X,P),tipo_valor(T),P==pensionado_administrativo,T==tipo_B,salario_minimo_conven(J,U),S is J*2.

      valor_incremento_vigencia(0.07,2016).
      incremento_cotizacion(X,I):-valor_mesada(X,M),valor_incremento_vigencia(V,P), I is M*V.

      valor_ajuste(100000).
      numero_dias(17078900,10).
      valor_ajuste_mesada(cc,M,J):-valor_pension(X,Y),numero_dias(X,D), J is (Y/30)*D.

      valores(X,T,P,L):-findall((V),((factor(X,T,Y,N,Z,P),Y==porcentaje,N==pago_soli,aporte_fondoSoli(X,W),V is W)),L).
      valores(X,T,P,L):-findall((V),((factor(X,T,Y,N,Z,P),Y==porcentaje,N==pago_salud,aporte_salud(X,C),V is C)),L).
      valores(X,T,P,L):-findall((V),((factor(X,T,Y,N,Z,P),Y==fijo,V is Z)),L).

      factor(17078900,descuento,porcentaje,pago_soli,Z,2016).
      factor(17078900,descuento,porcentaje,pago_salud,Z,2016).
      factor(17078900,descuento,fijo,ajubilud,206653,2016).

      suma_elementos([],0).
      suma_elementos([X|Xs],S):-suma_elementos(Xs,S2), S is S2 + X.

      total_descuentos(X,L,V):-valores(X,T,P,L),suma_elementos(L,S), V is S.

      prueba_descuentos(X,L,V):-suma_elementos(L,S), V is S.

      pago_neto(X,D,E,G):-suma_elementos(D,S1), suma_elementos(E,S2), G is S1 - S2.


    `)
  /*categoria_sustituto:= m.ProveAll(`categoria_sustituto(X,C).`)
  for _, solution := range categoria_sustituto {subsidio_libros(X,S):-tipo_pensionado(X,PA),tipo_valor(TA,salario_minimo_convencional(J,P)),S is J.
  subsidio_libros(X,S):-tipo_pensionado(X,PA),tipo_valor(TB,salario_minimo_convencional(J,P)),S is J*2.  //Verificarsubsidio_libros(X,S):-tipo_pensionado(X,PA),tipo_valor(TA,salario_minimo_convencional(J,P)),S is J.
  subsidio_libros(X,S):-tipo_pensionado(X,PA),tipo_valor(TB,salario_minimo_convencional(J,P)),S is J*2.  //Verificar
      //fmt.Printf("%s valor_pension -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
      fmt.Printf("%s\n",solution.ByName_("C"))
  }*/
  /*pago_aporte_solidaridad:= m.ProveAll(`aporte_fondoSoli(17078900,W).`)
  for _, solution := range pago_aporte_solidaridad{
    fmt.Printf("%s\n",solution.ByName_("W"))
  }*/
  /*aporte_salud:= m.ProveAll(`aporte_salud(17078900,C).`)
  for _, solution:=range aporte_salud {
    fmt.Printf("%s\n",solution.ByName_("C"))
  }*/
  /*subsidio_familiar:= m.ProveAll(`subsidio_familiar(17078900,Y).`)
  for _, solution := range subsidio_familiar{
    fmt.Printf("%s\n",solution.ByName_("Y"))
  }*/
  /*pago_subfamiliar := m.ProveAll(`pago_subfamiliar(17078900,F).`)
  for _, solution:= range pago_subfamiliar{
    fmt.Printf("%s\n",solution.ByName_("F"))
    }*/
  /*subsidio_libros:=m.ProveAll(`pago_subsidio_libros(17078900,S).`)
  for _, solution:= range subsidio_libros{
    fmt.Printf("%s\n",solution.ByName_("S"))
  }*/
  /*incremento_cotizacion_salud:= m.ProveAll(`incremento_cotizacion(17078900,I).`)
  for _, solution:= range incremento_cotizacion_salud{
    fmt.Printf("%s\n ", solution.ByName_("I"))
  }*/
  /*descuento:= m.ProveAll(`valores(17078900,descuento,2016,L).`)
  for _, solution:= range descuento{
    fmt.Printf("%s\n ", solution.ByName_("L"))
  }*/
  /*lista_descuento:= m.ProveAll(`valores(17078900,descuento,2016,L).`)
  for _, solution:= range lista_descuento{
    fmt.Printf("%s\n ", solution.ByName_("L"))
  }*/
  /*total_descuento:= m.ProveAll(`total_descuentos(17078900,V).`)
  for _, solution:= range total_descuento{
    fmt.Printf("result: %s\n ", solution.ByName_("V"))
  }*/
  /*total_descuento:= m.ProveAll(`prueba_descuentos(17078900,[317920,1907569,206653],V).`)
  for _, solution:= range total_descuento{
    fmt.Printf("valor descuento: %s\n ", solution.ByName_("V"))
  }*/

  pago_neto:= m.ProveAll(`pago_neto(X,[15896409],[317920,1907569,206653],G).`)
  for _, solution:= range pago_neto{
    fmt.Printf("pago neto: %s\n ", solution.ByName_("G"))
  }

}
