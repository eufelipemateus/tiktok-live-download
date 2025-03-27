# 🎥 TikTok Live Download  

Este é um projeto open source para processar lives do TikTok automaticamente, transformando-as em vídeos finais prontos para publicação. Ele utiliza uma biblioteca de pós-processamento de vídeo para gerar o vídeo final, disponível em: [TikTok Live PosProcessor](https://github.com/eufelipemateus/tiktok-live-posprocessor).  

## 🔗 Como funciona?  

1. A aplicação obtém o chat e a quantidade de viewers da live utilizando a biblioteca: [go-tiktoklive](https://github.com/steampoweredtaco/gotiktoklive).  
2. O software processa as lives e aplica efeitos com pós-processamento.  
3. Para rodar o software e baixar uma live, basta renomear o arquivo `lives.txt.example` para `lives.txt` e substituir `usuario` pelo nome do usuário alvo.  

## 🚀 Suporte a múltiplos usuários  

A aplicação pode processar múltiplos usuários simultaneamente. **No entanto, isso não é recomendado**, pois pode causar **rate limits** no TikTok devido ao grande número de requisições simultâneas.  

## 🌍 Rodando no Brasil  

O TikTok bloqueou o acesso ao seu app no Brasil, então, para rodar essa aplicação no país, é necessário utilizar um proxy. No Docker, há um serviço Tor embutido que permite **bypassar essa limitação automaticamente**.  
