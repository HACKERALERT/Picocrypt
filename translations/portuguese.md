<p>English | <a href="/translations/french.md">Français</a> | <a href="/translations/spanish.md">Español</a> | <a href="/translations/german.md">Deutsch</a> | <a href="/translations/portuguese.md">Português</a> | <a href="/translations/turkish.md">Türkçe</a> | <a href="/translations/chinese.md">中文</a> | <a href="/translations/russian.md">русский</a></p>
<p align="center"><img align="center" src="/images/Picocrypt.svg" width="512" alt="Picocrypt"></p> 

Picocrypt é uma ferramenta de criptografia muito pequena (daí <i>Pico</i>), muito simples mas muito segura, que você pode usar para proteger seus arquivos, gerar checksums, triturar arquivos, e muito mais. Ele foi projetado para ser a ferramenta <i>número um</i> de criptografia, com foco em segurança, simplicidade e confiabilidade. Picocrypt usa a cifra segura XChaCha20 e a função hash SHA3 para fornecer um alto nível de segurança, até mesmo de "agências de três letras" como a NSA. Ele foi projetado para máxima segurança, sem comprometê-la em absolutamente nada, e é construído com os módulos x/crypto padrões do Go. <strong>Sua privacidade e segurança estão sob ataque. Recupere-as com confiança protegendo seus arquivos com Picocrypt.</strong>

<p align="center"><img align="center" src="/images/Picocrypt.png" width="384" alt="Picocrypt"></p>

# Financiamento
Por favor, doe para o Picocrypt no <a href="https://opencollective.com/picocrypt">Open Collective</a> para arrecadar dinheiro para uma potencial auditoria da Cure53. Como este é um projeto no qual eu gasto muitas horas e não me gera dinheiro, não posso eu mesmo pagar por uma auditoria. <i>Picocrypt precisa do apoio de sua comunidade.</i>

# Downloads
<strong>Importante:</strong> Há um abandonware desatualizado e inútil chamada PicoCrypt na Internet, que foi atualizado pela última vez em 2005. O PicoCrypt não está relacionado de forma alguma ao Picocrypt (este projeto). Certifique-se de baixar o Picocrypt apenas do site oficial ou deste repositório para garantir que você receba o Picocrypt autêntico e sem backdoors.

## Windows
Picocrypt para Windows é extremamente simple. Para baixar o executável portátil, autônomo e mais recente para Windows, clique <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.16/Picocrypt.exe">aqui</a>. Se o Windows Defender ou seu antivírus sinalizar Picocrypt como um vírus, por favor, faça sua parte e registre como um falso positivo, para a melhoria de todos.

## macOS
Picocrypt para macOS também é muito simples. Baixe o Picocrypt <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.16/Picocrypt.app.zip">aqui</a>, extraia o arquivo zip e execute o Picocrypt que está dentro. Se você não conseguir abrir o Picocrypt porque não é de um desenvolvedor verificado, clique com o botão direito do mouse no Picocrypt e clique em "Abrir". Se você ainda receber o aviso, clique com o botão direito do mouse no Picocrypt e clique novamente em "Abrir" e você deve conseguir iniciar o Picocrypt.

## Linux
Um Snap está disponível para Linux. Certifique-se de ter instalado o [Snapcraft](https://snapcraft.io/) (`sudo apt install snapd`) e instale o Picocrypt: `sudo snap install picocrypt`. Devido à complexidade das dependências e de links estáticos, eu não distribuo binários autônomos .deb ou .rpm porque eles não seriam confiáveis e não valeriam a pena. O Snapcraft gerencia todas as dependências automaticamente e é a forma recomendada de executar o Picocrypt em qualquer grande distribuição do Linux.

# Por que Picocrypt?
Por que você deve usar Picocrypt em vez de BitLocker, NordLocker, VeraCrypt, AxCrypt, ou 7-Zip? Aqui estão algumas razões pelas quais você deve escolher Picocrypt:
<ul>
	<li>Ao contrário de NordLocker, BitLocker, AxCrypt e a maioria dos provedores de armazenamento em nuvem, Picocrypt e suas dependências são completamente auditáveis e de código aberto. Você pode verificar por si mesmo que não há backdoors ou falhas.</li>
	<li>Picocrypt é <i>minúsculo</i>. Enquanto NordLocker tem mais de 100MB e VeraCrypt tem mais de 30MB, Picocrypt tem apenas 3MB, aproximadamente o tamanho de uma imagem de alta resolução. E isso não é tudo - Picocrypt é portátil (não precisa ser instalado) e não requer privilégios de administrador/root.</li>
	<li>Picocrypt é mais fácil e mais produtivo de usar do que o VeraCrypt. Para criptografar arquivos com VeraCrypt, você teria que gastar pelo menos cinco minutos configurando um volume. Com a interface simples do Picocrypt, tudo o que você precisa fazer é arrastar e soltar seus arquivos, digitar uma senha e pressionar Iniciar. Todos os procedimentos complexos são tratados internamente pelo Picocrypt. Quem disse que criptografia segura não pode ser simples?</li>
	<li>Picocrypt é projetado para segurança. O 7-Zip é uma ferramenta de arquivamento, não de criptografia; portanto, seu foco não é a segurança. Já o Picocrypt é construído com a segurança como prioridade número um. Cada parte do Picocrypt existe por uma razão, e qualquer coisa que possa impactar a segurança do Picocrypt é removida. O Picocrypt é construído com criptografia em que você pode confiar.</li>
	<li>Picocrypt também autentica os dados além de protegê-los, impedindo que hackers modifiquem dados sensíveis maliciosamente. Isto é útil quando você está enviando arquivos criptografados através de um canal inseguro e quer ter certeza de que eles cheguem intactos. Picocrypt usa HMAC-SHA3 para autenticidade, que é uma função de hash altamente segura em uma construção bem conhecida.</li>
	<li>Picocrypt protege ativamente os dados criptografados do cabeçalho contra corrupção adicionando bytes extras de paridade Reed-Solomon, então se parte do cabeçalho de um volume (que contém componentes criptográficos importantes) se corromper (por exemplo, degradação de bits no disco rígido), Picocrypt ainda pode recuperar o cabeçalho e desgriptografar seus dados com uma alta taxa de sucesso. Picocrypt também pode codificar o volume inteiro com Reed-Solomon para evitar qualquer corrupção em seus arquivos importantes.</li>
</ul>

Ainda não se convenceu? Veja abaixo mais razões ainda por que o Picocrypt se destaca do resto.

# Recursos
Picocrypt é uma ferramenta muito simples, e a maioria dos usuários entenderá intuitivamente como utilizá-lo em poucos segundos. Em um nível básico, simplesmente soltar seus arquivos, digitar uma senha e pressionar Iniciar é tudo o que é necessário para criptografar seus arquivos. Bem simples, não é?

Embora seja simples, Picocrypt também almeja ser poderoso nas mãos de usuários experientes e avançados. Assim, há algumas opções adicionais que você pode usar para atender às suas necessidades.
<ul>
	<li><strong>Gerador de senhas</strong>: Picocrypt fornece um gerador de senhas seguras que você pode usar para criar senhas criptograficamente seguras. Você pode personalizar o tamanho da senha, assim como os tipos de caracteres a serem incluídos.</li>
	<li><strong>Metadados do arquivo</strong>: arquivo: Use isto para armazenar notas, informações e texto junto com o arquivo (eles não serão criptografados). Por exemplo, você pode colocar uma descrição do arquivo que você está criptografando antes de enviá-lo a alguém. Quando a pessoa a quem você enviou soltar o arquivo no Picocrypt, sua descrição será mostrada a essa pessoa.</li>
	<li><strong>Modo rápido</strong>: Usar este modo irá acelerar muito a encriptação/desencriptação. Neste modo, o BLAKE2b será usado para autenticar os dados em vez do SHA3. Isto proporciona maiores velocidades, mas com uma margem de segurança ligeiramente menor.</li>
	<li><strong>Modo paranoico</strong>: Usar este modo irá criptografar seus dados com ambos XChaCha20 e Serpent, em cascata. Isto é recomendado para proteger arquivos ultra-secretos e proporciona o mais alto nível de segurança praticamente atingível. Para que um hacker acesse seus dados criptografados, tanto a cifra XChaCha20 quanto a Serpent devem ser quebradas, presumindo que você tenha escolhido uma boa senha.</li>
	<li><strong>Prevenir corrupção usando Reed-Solomon</strong>: Este recurso é muito útil se você estiver planejando arquivar dados importantes em um servidor na nuvem ou mídia externa por um longo tempo. Se a opção for marcada, o Picocrypt usará o código de correção de erros Reed-Solomon para adicionar 8 bytes extras para cada 128 bytes para evitar a corrupção do arquivo. Isto significa que até ~3% de seu arquivo pode corromper e o Picocrypt ainda será capaz de corrigir os erros e descriptografar seus arquivos sem corrupção. É claro que se seu arquivo se corromper demais (por exemplo, se você derrubar seu disco rígido), Picocrypt não será capaz de recuperar totalmente seus arquivos, mas tentará o melhor para recuperar o que puder. Note que esta opção irá tornar a encriptação/desencriptação consideravelmente mais lentas.</li>
	<li><strong>Manter a saída descriptografada mesmo que esteja corrompida ou modificada</strong>: Picocrypt checa automaticamente a integridade após a desencriptação. Se o arquivo foi modificado ou está corrompido, Picocrypt apagará automaticamente a saída para a segurança do usuário. Se você quiser manter os dados corrompidos ou modificados após a desencriptação, marque esta opção. Além disso, se esta opção for marcada e o recurso Reed-Solomon tiver sido usado no arquivo criptografado, Picocrypt tentará recuperar o máximo possível do arquivo durante a desencriptação.</li>
</ul>

Além destas opções de encriptação/desencriptação, Picocrypt também fornece um triturador seguro de arquivos e um gerador de checksum.

# Segurança
Para mais informações sobre como Picocrypt lida com a criptografia, consulte os detalhes técnicos em <a href="Internals.md">Internals</a> Se você está preocupado com a minha segurança ou a deste projeto, deixe-me assegurar-lhe que este repositório não será roubado ou terá backdoors introduzidos. Tenho 2FA / TOTP (autenticação de dois fatores) habilitada em todas as contas ligadas ao Picocrypt (GitHub, Google, Reddit, Discord etc), além de criptografia de disco completo em todos os meus dispositivos portáteis. Para mais reforço ainda, Picocrypt usa meus forks "offline" de dependências, e eu só pego o upstream depois de dar uma olhada nas mudanças e acreditar que não há nenhum problema de segurança. Você pode usar o Picocrypt com confiança.

# Comunidade
Aqui estão alguns lugares onde você pode ficar em dia com Picocrypt e se envolver:
<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>
Eu recomendo fortemente que você se inscreva no subreddit do Picocrypt, porque todas as atualizações e enquetes serão postadas lá. Lembre-se de confiar somente nestas redes sociais e tenha cuidado com hackers que possam tentar se fazer passar por mim. Nunca lhe pedirei sua senha, então se alguém pedir, não será eu. Eu nunca lhe direi para baixar um arquivo de um link suspeito, e qualquer um que diga não será eu.

# Stargazers
Como está o Picocrypt? Dê uma olhada abaixo para descobrir.
[![Stargazers over time](https://starchart.cc/HACKERALERT/Picocrypt.svg)](https://starchart.cc/HACKERALERT/Picocrypt)

# Contribuições
Se você encontrar um bug ou tiver uma solicitação de recurso, por favor, entre em contato comigo ou crie um Issue. Estou aberto a colaboradores e todos os tipos de contribuições são bem-vindas. Se você encontrar um problema de segurança, por favor, siga as instruções aqui para relatá-lo.

Se você é multilíngue e conhece um idioma para o qual o Picocrypt ainda não foi traduzido, eu adoraria ter sua ajuda para traduzir esta página, assim como a interface do Picocrypt. Quanto mais idiomas, mais acolhedor!

# Doações

Se você achar útil o Picocrypt, por favor, considere me apoiar no <a href="https://patreon.com/evansu">Patreon</a>. Estou fornecendo este software completamente de graça, e adoraria ter alguns apoiadores que me motivarão a continuar meu trabalho no Picocrypt.

# Thank You's
Um obrigado do fundo do meu coração aos meus apoiadores no Patreon:
<ul>
	<li>Frederick Doe</li>
</ul>

E às pessoas no Open Collective que fizeram contribuições significativas:
<ul>
	<li>jp26 ($50)</li>
</ul>
Vocês são as pessoas que me inspiram a trabalhar no Picocrypt e o fornecem gratuitamente a todos!

Também um enorme obrigado à seguinte lista de cinco pessoas, que foram as primeiras a doar e apoiar o Picocrypt:
<ul>
	<li>W.Graham</li>
	<li>N. Chin</li>
	<li>Manjot</li>
	<li>Phil P.</li>
	<li>E. Zahard</li>
</ul>

Além disso, um grande obrigado a essas pessoas, que ajudaram a traduzir Picocrypt e torná-lo mais acessível ao mundo:
<ul>
	<li>@umitseyhan75 para turco</li>
	<li>@digitalblossom para o alemão</li>
	<li>@zeeaall para o português brasileiro</li>
</ul>

Finalmente, agradecimentos a essas pessoas por me ajudarem quando necessário:
<ul>
	<li>Fuderal no Discord por me ajudar a configurar um servidor Discord</li>
	<li>u/greenreddits pelo feedback e apoio constantes</li>
	<li>u/Tall_Escape por me ajudar a testar o Picocrypt</li>
	<li>u/NSABackdoors por fazer muitos testes</li>
</ul>
