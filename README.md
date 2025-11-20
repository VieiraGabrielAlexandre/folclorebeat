# 🐺 **FolcloreBeat – Beat ’em Up Brasileiro com Ebitengine**

**Um jogo em Go que mistura ação estilo Streets of Rage + transformação à la Altered Beast, com bosses lendários do folclore brasileiro.**

---

## 🎮 **O que é este jogo?**

**FolcloreBeat** é um jogo *beat ’em up* desenvolvido em **Go** usando o **Ebitengine**, voltado para homenagear personagens clássicos do folclore brasileiro — como **Saci**, **Cuca** e **Mula sem Cabeça** — em uma aventura brutal, divertida e cheia de habilidades especiais.

Você controla um herói que, ao derrotar criaturas sombrias, coleta essências místicas e desbloqueia sua verdadeira forma:
🩸 **O Lobisomem**, com golpes mais fortes e velocidade aprimorada.

---

# 🕹️ **Sobre o Ebitengine**

O jogo usa o **[Ebitengine](https://ebitengine.org/)** (antigo Ebiten), principal engine 2D para Go.

### **📌 Por que usar o Ebitengine?**

* É simples, leve e multiplataforma
* 100% feito para jogos 2D
* Suporta:

    * áudio
    * sprites
    * input (teclado, joystick)
    * física simples
    * renderização acelerada
* Faz build para:

    * Linux
    * Windows
    * macOS
    * Web (WebAssembly)
    * Mobile (Android/iOS)

### **📌 Quando usar Ebitengine?**

* Jogos 2D, pixel art ou top-down
* Prototipagem rápida
* Projetos que precisam buildar fácil
* Engines leves para projetos solo/indie
* Go + GameDev (ótima combinação)

### **📌 Como instalamos**

No projeto:

```bash
go get github.com/hajimehoshi/ebiten/v2
```

### **📌 Documentação oficial**

* [https://ebitengine.org/en/documents/](https://ebitengine.org/en/documents/)
* Exemplos: [https://github.com/hajimehoshi/ebiten/tree/main/examples](https://github.com/hajimehoshi/ebiten/tree/main/examples)
* Tutoriais: [https://ebitengine.org/en/documents/guide/start.html](https://ebitengine.org/en/documents/guide/start.html)

---

# 🧠 **Lógica do Desenvolvimento**

O jogo segue uma arquitetura limpa e legível, dividida por responsabilidade:

```
folclorebeat/
  cmd/game           → executável do jogo
  internal/
    engine/          → game loop, HUD, controle global
    player/          → player, ataques, física, transformação
    enemies/         → zumbis e vampiros
    bosses/          → Saci, Cuca (e os próximos)
    powerups/        → orbes, upgrades
    world/           → mundo, fase, cenário
    combat/          → hitbox/hurtbox lógica
```

### **Game loop do Ebitengine**

O cycle padrão:

```
Update() → game logic (movimento, IA, combate)
Draw()   → renderiza sprites/retângulos
Layout() → tamanho lógico da tela
```

### **Player Lógica**

* Anda
* Pula
* Soco (A)
* Chute (S)
* Voadora (S no ar)
* Recebe dano com i-frames
* Recebe XP e se transforma em lobisomem

### **Enemy Lógica**

* Zumbi e Vampiro caminham até o player
* Ao morrer, dropam **orbes**
* Orbes dão XP ao serem coletados

### **Boss Lógica**

#### ✔️ Saci (Boss 1)

* Teleporta ao redor do player
* Dá dano por contato
* HP médio
* Morreu → XP extra

#### ✔️ Cuca (Boss 2)

* Se move pela fase
* Atira **fireballs diagonais** no player
* HP alto
* Morreu → abre caminho para o próximo boss

---

# 📚 **História do Jogo (Lore)**

No coração de uma noite sem lua, algo sombrio desperta no Brasil profundo.

Os antigos espíritos da mata ruíram, criaturas mortas-vivas vagam pelos caminhos… E, nas sombras, forças esquecidas querem destruir o equilíbrio entre o mundo dos vivos e dos mitos.

Você é **Alexandre**, um jovem amaldiçoado, vítima de uma linhagem ancestral que carrega o sangue do **Lobisomem**. A cada inimigo derrotado, sua alma absorve fragmentos espirituais que despertam seu poder interior.

Para restaurar a ordem e salvar a humanidade, Alexandre deve enfrentar:

* hordas de **zumbis** e **vampiros**;
* o **Saci**, mestre das ilusões e teletransportes;
* a **Cuca**, bruxa reptiliana capaz de lançar bolas de fogo;
* e, futuramente, a **Mula sem Cabeça**, avatar flamejante da fúria.

Somente ao dominar sua metamorfose lupina, ele poderá derrotar os monstros e reequilibrar o folclore.

---

# ✨ **O que já foi implementado**

### ✔️ Player

* Movimento lateral
* Pulo
* Soco
* Chute
* Voadora
* Transformação em **Lobisomem**
* Barra de HP e XP

### ✔️ Inimigos

* Zumbi (IA simples)
* Vampiro (IA rápida)

### ✔️ Powerups

* Orbe místico flutuante
* Coletado ao tocar
* Dá XP
* Faz o player evoluir

### ✔️ Boss 1: Saci

* Teleporte
* Contato causa dano
* HP + comportamento único

### ✔️ Boss 2: Cuca

* Movimentação horizontal
* Fireballs **diagonais**
* HP alto

---

# 🔜 **Próximos passos (Roadmap)**

### 🐎 Boss 3 – Mula Sem Cabeça

* Dash flamejante
* Rastro de fogo no chão
* Explosão quando fica com pouca vida
* Grito que stunna o player

### 🗺️ Fase Completa

* Tileset de cemitério
* Tileset de mata
* Parallax background
* Spawner de ondas de inimigos

### 🎨 Gráficos

* Sprites animados
* Efeitos (hitflash, explosão, partículas)
* HUD estilizado

### 🔊 Áudio

* Trilhas
* Efeitos (passo, golpe, dano, transformação)

---

# 💻 **Como rodar**

### Pré-requisitos (Linux)

```bash
sudo apt install \
  libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev \
  libxi-dev libasound2-dev libglu1-mesa-dev libxxf86vm-dev
```

### Rodar o jogo

```
go run ./cmd/game
```

### Buildar

```
go build -o folclorebeat ./cmd/game
```

---

# 📦 **Estrutura de pastas**

```
folclorebeat/
├── cmd/
│   └── game/
│       └── main.go
├── internal/
│   ├── engine/
│   │   ├── game.go
│   │   ├── hud.go
│   │   └── ...
│   ├── player/
│   ├── enemies/
│   ├── bosses/
│   ├── powerups/
│   ├── combat/
│   ├── world/
│   └── ...
└── go.mod
```

---

# 🚀 **Por que este projeto é foda?**

* **É um beat ’em up completo feito em Go**
* Mistura programação funcional e criativa
* Usa folclore brasileiro de forma divertida
* Cresce organicamente com novas fases e bosses
* É perfeito para aprender:

    * game loop
    * colisões
    * IA simples
    * física
    * sprites
    * estados
    * arquitetura de jogos

---

# 🙌 Contribuições

Sinta-se livre para abrir issues, discutir ideias, sugerir personagens novos ou movimentos do lobisomem.

E claro…
se quiser deixar sua marca no folclore digital, as portas estão abertas. 🐺🇧🇷

---
