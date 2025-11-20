# 📄 **Dossiê Técnico – FolcloreBeat**

### *Desenvolvimento, Arquitetura e Decisões Técnicas do Jogo*

---

# 🧱 1. Arquitetura Geral do Projeto

O projeto segue um padrão simples, porém organizado, inspirado em:

* Clean Architecture
* Estrutura modular por domínio
* Componentização funcional

```
folclorebeat/
  cmd/game/           → entrada da aplicação (main.go)
  internal/engine/    → game loop, efeitos globais, HUD
  internal/player/     → protagonista, física, ataques, transformação
  internal/enemies/    → inimigos básicos (zumbi, vampiro)
  internal/bosses/     → bosses com comportamentos avançados
  internal/powerups/   → orbes e upgrades
  internal/combat/     → sistema de hitbox e colisões
  internal/world/      → estágio, cenário
```

Cada módulo é responsável por uma única parte do jogo, evitando acoplamento.

---

# 🎮 2. O Ebitengine e o Game Loop

Todos os jogos no Ebitengine seguem este padrão:

### **Update()**

É a lógica do jogo chamada ~60 vezes por segundo.

Aqui colocamos:

* Movimento do player
* IA do inimigo
* Ataques
* Colisões
* Atualização de power-ups
* Avanço para bosses
* Consumo de input
* Física (gravidade, pulo)

### **Draw()**

Desenha tudo na tela.

Hoje usamos placeholders (retângulos), mas já preparado para sprites.

### **Layout()**

Define a resolução **lógica** (480×270), independente da janela real.

Permite pixel art consistente.

---

# 🧍 3. O Player: Estados, Física e Combate

O player tem vários estados:

```
Idle       → parado
Walk       → andando
Jump       → pulando
Punch      → soco
Kick       → chute
AirKick    → voadora
Wolf       → Lobisomem (versão evoluída)
```

Esses estados orientam:

* lógica interna do player
* hitbox de ataque
* animações futuras
* limitações (ex: não andar durante soco)

---

## 🏃 3.1 Movimento

O movimento é simples, usando `X` e `Y`:

```go
if KeyRight → X += 2
if KeyLeft  → X -= 2
```

O player sempre mantém um **facing** (1 direita / -1 esquerda) para saber para onde atacar.

---

## 🦘 3.2 Pulo (física)

Implementamos uma “física minimalista”:

* **VY** é a velocidade vertical
* Quando pula, VY = -6
* A cada frame, adicionamos gravidade `VY += 0.25`
* Y += VY

Quando Y atinge o chão (200), o player para:

```go
if Y >= 200 → OnGround = true, VY = 0
```

Simples, eficiente, estilo beat ’em up clássico.

---

# 👊 3.3 Sistema de Ataque do Player

Ataques são controlados por estado e cooldown:

### **Soco (A)**

* curto alcance
* rápido cooldown

### **Chute (S)**

* mais alcance
* cooldown maior

### **Voadora (S no ar)**

* ataque aéreo
* dano eficiente contra bosses

Esse mecanismo cria variedade sem complexidade.

---

# 📏 3.4 Hitbox de Ataque (AttackHitbox)

O ataque cria uma pequena área na frente do player:

```go
if player facing right  → hitbox nasce ao lado direito  
if player facing left   → hitbox nasce ao lado esquerdo  
```

Essa hitbox é um retângulo:

```go
X, Y, W, H
```

Que é comparado com a hitbox do inimigo:

```go
if atkRect.Intersects(enemyRect) → dano!
```

---

# 🧪 4. Sistema de Colisão: combat.Rect

Criamos uma estrutura genérica:

```go
type Rect struct { X, Y, W, H float64 }
func (r Rect) Intersects(o Rect) bool
```

Isso permite:

* colisão de ataques
* colisão de fireballs
* colisão com boss
* coleta de orbes

Solução pequena, eficiente e universal no jogo.

---

# 💀 5. Inimigos: Zumbi e Vampiro

Eles têm:

```
X, Y
VX → velocidade
HP
Hitbox
Alive/Killed
XPReward
```

### IA simples:

Mover na direção do player:

```go
dx := player.X - e.X
if dx > 0 → X += VX
if dx < 0 → X -= VX
```

Isso recria o comportamento “walk forward” dos beat ’em ups antigos.

---

# 🧿 6. Power-ups (Orbes)

Quando o inimigo morre:

* criamos um orb `powerups.NewWolfOrb()`
* ele flutua com uma senóide (`sin`)
* se player toca → coleta

### Funcionalidade:

* aumenta XP
* eventualmente transforma o player em Lobisomem

---

# 🐺 7. Transformação em Lobisomem

Implementada no player:

```go
func TransformToWolf() {
    IsWolf = true
    AttackPower = 3
}
```

O player ganha:

* força maior
* estética diferente
* próprio estado animável

Inspirado diretamente em **Altered Beast**.

---

# 🎨 8. HUD (Health e XP Bars)

HUD desenhado em:

`engine/hud.go`

Usa:

* barra de fundo
* barra preenchida proporcional ao valor

HP = verde
XP = azul

Fácil de estilizar depois.

---

# 👹 9. Boss 1 – SACI: lógica completa

Comportamentos:

### ✨ Teleporte

A cada ~1.5s escolhe nova posição perto do player:

```go
dir := ±1
dist := 40–120
X = player.X + dir*dist
```

Controlado por cooldown.

### ✨ Flutuação

Adiciona vida visual:

```
Y = baseY + sin(frame * 0.1) * 4
```

### ✨ Contato causa dano

Ao encostar:

```go
player.TakeDamage(1)
```

### ✨ Morte

Desbloqueia boss 2 (Cuca).

---

# 🐊🔥 10. Boss 2 – CUCA: lógica completa

Cuca tem um sistema mais avançado:

---

## 10.1 Movimentação horizontal

Ela segue o player lentamente:

```
if player está à direita → X += 0.6
se player está à esquerda → X -= 0.6
```

---

## 10.2 Fireballs **diagonais**

Essa foi a parte mais legal:

* Cuca calcula um vetor da posição dela até o player:

```go
dx = player.X - cuca.X
dy = player.Y - cuca.Y
dist = sqrt(dx² + dy²)
dx /= dist
dy /= dist
```

* Normaliza para obter direção
* Multiplica pelo speed

Fireballs agora vão:

👉 **direção do player**
👉 **em diagonal**, não horizontal

Isso cria um mini “bullet hell”.

---

## 10.3 Fireball física

Fireballs usam:

```
X += VX
Y += VY
```

e têm:

* velocidade variável
* morte ao sair da tela
* colisão com player

---

## 10.4 Cuca flutua igual ao Saci

```
Y = baseY + sin(frame * 0.07) * 5
```

---

# 🔥 11. Engine de Fase e Progressão

`Game.Update()` controla progressão:

1. Mata inimigos comuns → spawn do Saci
2. Mata Saci → XP extra + spawn da Cuca
3. Mata Cuca → deixa preparado para próximo boss

Isso cria um **loop de gameplay** sólido e expansível.

---

# 🐎 12. Próximo Boss: Mula sem Cabeça (pré-planejado)

Algoritmo sugerido:

* Dashes rápidos na direção do player
* Rastro de fogo persistente (área de dano)
* “Explosão” ao chegar em HP crítico
* Grito que stunna o player

Será mais avançado que Cuca.

---

# 🛠️ 13. Por que cada decisão foi tomada?

### ✔️ Ebitengine →

Porque queremos algo leve, Go-nativo e fácil de rodar em todas as plataformas.

### ✔️ Estrutura modular →

Permite evoluir o jogo sem virar bagunça.

### ✔️ Física própria →

Beat ’em ups não precisam de física realista.
Só:

* gravidade
* salto
* colisões retangulares

Perfeito para este estilo.

### ✔️ Retângulos primeiro →

Prototipar rápido, testar IA, lógica, sensação antes de arte.

### ✔️ IA simples dos inimigos →

Beat ’em ups clássicos usam IA simples + quantidade.

### ✔️ Bosses cada um com arquivo próprio →

Cada boss é complexo o suficiente para merecer sua lógica isolada.

### ✔️ Power-ups flutuantes →

Feedback visual direto, estilo arcade.

### ✔️ HUD minimalista →

Clareza antes do estilo.

---

# 📌 14. O que falta para o jogo virar “jogo de verdade”

* sprites animados (estou pensando em usar LibreSprite, Aseprite, Piskel)
* tileset para cenários
* efeitos visuais (hitflash, partículas)
* músicas/sons
* cutscenes
* sistema de menus
* combinações de golpes (combo system)