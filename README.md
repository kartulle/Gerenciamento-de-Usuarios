# Gerenciamento-de-Usuarios - Documentação
A API de gerenciamento de usuários foi programada na linguagem GOLANG com o intuito de gerenciar os usuários de um banco online. A api é limpa e eficiente, e graças à linguagem utilizada, seus mecanismos permitem a construção de um programa flexível e modular.
</br>
</br>
<h1>Primeiros Passos:</h1>
1. Baixe o Banco de dados postgresql. No pacote do site oficial, virá instalado o pgadmin4;</br>
2. Abra o pgadmin4 e digite a senha cadastrada no primeiro acesso;</br>
3. Para usar o nosso banco de dados, será necessário importar o arquivo disponibilizado como “user_management”. Para isso, será necessário configurar o path do pgadmin: file>preferences>paths>binary paths>;</br>
4. Procure a opção PostgreSQL 16 e cole o caminho do diretório da pasta Bin;</br>
5. Agora crie um database com nome “user_management”;</br>
6. Clique com o botão direito do mouse em cima dele e, em seguida, clique em “Restore”;</br>
7. No input de filename, selecione o ícone de pastinha e escolha o arquivo disponibilizado de nome “user_management” mencionado anteriormente;</br>
8. Clique em Restore.</br>
</br>
Agora que você já tem o banco de dados, vamos configurar o código ao banco de dados. Procure o diretório Database, entre no arquivo db.go e coloque a senha cadastrada no seu banco de dados, verifique também se a porta e o nome estão corretos para evitar falhas de funcionamento. Finalmente, podemos rodar a nossa aplicação com o comando “go run main.go”. Aproveite!</br>

</br>
<h1> Principais funcionalidades: </h1>

<h2>Empregado:</h2>
1.1 Inserir usuário no banco; </br>
1.2 Editar usuário;</br>
1.3 Remover usuário do banco;</br>
1.4 Listar usuários;</br>
1.5 Ver ID da transação bancária.</br>

<h2>Cliente:</h2>
1.1 Cadastrar no banco; </br>
1.2 Editar suas informações; </br>
1.3 Remover Conta; </br>
1.4 Ver ID da transação bancária; </br>
