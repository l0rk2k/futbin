# Guía de Uso de Futbin CLI

Esta herramienta te permite acceder a la base de datos de Futbin desde tu terminal. Aquí encontrarás ejemplos detallados de cómo utilizar los distintos comandos y filtros.

## Comandos Principales

### 1. Ver Jugadores Recientes
Muestra los últimos jugadores agregados a la base de datos.
```bash
./futbin latest
```

### 2. Ver Jugadores Populares
Muestra los jugadores más buscados en este momento.
```bash
./futbin popular
```

### 3. Buscar Jugadores (`players`)
Este es el comando más potente, permite filtrar por múltiples criterios.

#### Filtros Básicos
- **Por Nombre/Versión:**
  ```bash
  # Buscar versiones "OTW" (Ones to Watch)
  ./futbin players --version otw
  ```

- **Por Media (Rating):**
  ```bash
  # Jugadores con media 90 o superior
  ./futbin players --ovr 90-
  
  # Jugadores con media entre 85 y 88
  ./futbin players --ovr 85-88
  ```

- **Por Precio:**
  ```bash
  # Jugadores que cuesten menos de 200,000 monedas
  ./futbin players --price -200000
  
  # Jugadores entre 10k y 50k
  ./futbin players --price 10000-50000
  ```

#### Filtros Avanzados (Nuevos)

- **PlayStyles (Estilos de Juego):**
  Busca jugadores que tengan un estilo de juego específico.
  ```bash
  # Buscar jugadores con "Finesse Shot"
  ./futbin players --playstyles "Finesse Shot"
  
  # Buscar jugadores con "Tiki Taka"
  ./futbin players --playstyles "Tiki Taka"
  ```

- **PlayStyles+:**
  Busca jugadores que tengan el estilo de juego en versión Plus (dorado).
  ```bash
  # Buscar jugadores con "Power Shot+"
  ./futbin players --playstyles-plus "Power Shot"
  ```

- **Tipo de Cuerpo (Body Type):**
  Filtra por el código de tipo de cuerpo.
  ```bash
  # Buscar jugadores con tipo de cuerpo específico
  ./futbin players --body-type "High & Average"
  ```

- **Likes:**
  Filtra por popularidad (número de likes en Futbin).
  ```bash
  # Jugadores con más de 1000 likes
  ./futbin players --likes 1000-
  ```

- **AcceleRATE:**
  Filtra por tipo de aceleración (solo Next Gen).
  ```bash
  # Buscar jugadores "Lengthy"
  ./futbin players --accelerate lengthy
  ```

#### Combinando Filtros
Puedes combinar todos los filtros que quieras para búsquedas específicas.

**Ejemplo: Brasileños en LaLiga con buen pase**
```bash
./futbin players --nation 54 --league 53 --passing 85-
```
*(Nation 54 = Brasil, League 53 = LaLiga)*

**Ejemplo: Delanteros rápidos y baratos**
```bash
./futbin players --position ST --pace 90- --price -50000
```

### 4. Consultar IDs de Ligas y Clubes
Para usar los filtros de liga y club, necesitas sus IDs numéricos.

- **Ver todas las ligas:**
  ```bash
  ./futbin leagues
  ```

- **Ver clubes de una liga específica:**
  ```bash
  # Ver clubes de la Premier League (ID 13)
  ./futbin clubs --league 13
  ```

## Consejos Adicionales
- Usa `--platform PC` o `--platform XB` si juegas en PC o Xbox (por defecto es PS).
- Puedes usar `head` para limitar los resultados: `./futbin latest | head -10`
