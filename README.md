# Worker Pool em Go

Estudo de processamento de eventos de jogo usando worker pool com goroutines,
channels e mutex.

## O que faz

- Cria uma fila de eventos (`Evento{Nome, Acao}`).
- Lança 3 workers concorrentes que leem do mesmo channel.
- Cada worker processa o evento e incrementa um contador compartilhado.
- Ao final, imprime o total de eventos processados.
