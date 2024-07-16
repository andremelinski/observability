[Collector](https://opentelemetry.io/docs/collector/)

-   The OpenTelemetry Collector offers a vendor-agnostic implementation of how to receive, process and export telemetry data. It removes the need to run, operate, and maintain multiple agents/collectors. This works with improved scalability and supports open source observability data formats (e.g. Jaeger, Prometheus, Fluent Bit, etc.) sending to one or more open source or commercial backends. The local Collector agent is the default location to which instrumentation libraries export their telemetry data

-   Coleta de dados de telemetria
-   colecta logs, metricas, etc.. comporta como um componente independente
-   Pode ser um agente ou um servico
-   Pode ser trabalhado no formato de pipeline: pegando o dado, tratando, convertendo e exportando pra algum sistema
-   Colector eh agnostico de vendor

Vendor

-   Quem consegue pegar as infos do colector,
-   Exemplo de vendor:Prometheus, cloudwatch, etc
