# GO-GIT-SCANNER

Este proyecto es un escáner de Git que realiza diferencias periódicas y las envía a un endpoint.

## Estructura del Proyecto

```
GO-GIT-SCANNER/
├── src/
│   └── utils/
│       ├── diff.go
│       ├── file.go
│       ├── post.go
│       └── scanner.go
├── Makefile
└── README.md
```

## Uso

1. Compilar el proyecto:

   ```
   make build
   ```

2. Ejecutar el escáner:

   ```
   make run
   ```

3. Limpiar los binarios:
   ```
   make clean
   ```

Asegúrate de configurar las constantes en los archivos (`gitDir`, `outputFile`, `apiEndpoint`) antes de ejecutar.

## Dependencias

Este proyecto utiliza:

- github.com/robfig/cron/v3

Asegúrate de ejecutar `go mod tidy` para gestionar las dependencias.

Hola mundo - 2024-08-22 21:33:00

Hola mundo - 2024-08-23 08:29:53

Hola mundo - 2024-08-23 08:34:44

Hola mundo - 2024-08-23 08:53:48

Hola mundo - 2024-08-23 08:57:58

Hola mundo - 2024-08-23 08:58:33

Hola mundo - 2024-08-23 09:01:39

Hola mundo - 2024-08-23 09:03:12

Hola mundo - 2024-08-23 09:08:03

Hola mundo - 2024-08-23 09:10:55

Hola mundo - 2024-08-23 09:27:06

Hola mundo - 2024-08-23 09:28:27

Hola mundo - 2024-08-23 09:36:51

Hola mundo - 2024-08-23 09:38:32

Hola mundo - 2024-08-23 09:39:06

Hola mundo - 2024-08-23 09:40:46

Hola mundo - 2024-08-23 16:07:41

Hola mundo - 2024-08-23 16:18:43

Hola mundo - 2024-08-23 16:20:15

Hola mundo - 2024-08-23 16:21:26

Hola mundo - 2024-08-23 16:27:49

Hola mundo - 2024-08-23 16:30:02

Hola mundo - 2024-08-23 17:26:20

Hola mundo - 2024-08-24 00:10:18
Resumen de las actividades - 2024-08-24 00:52:09
 **Resumen de las Actividades del Día Anterior:**

**Título: Definir cuenta interbancaria en una transferencia después de la asignación**
- **Contexto:** En esta actividad, se trabajó en el establecimiento y definición clara de cuentas bancarias para transferencias, con el propósito de asegurar un flujo eficiente de fondos entre diferentes instituciones financieras. Esto implicó la implementación de protocolos específicos que faciliten la identificación y asignación correcta de las cuentas en sistemas de pago electrónico, garantizando así una gestión precisa y oportuna de transacciones interbancarias.
- **Objetivo:** El objetivo principal fue asegurar que cada transferencia se realice a la cuenta bancaria correspondiente después de su asignación automatizada o manual, mediante el uso de tecnologías y procedimientos estandarizados que minimicen los riesgos asociados con errores en la identificación de cuentas.
- **Pasos Realizados:** Se revisaron y actualizaron las políticas y procedimientos relacionados con transferencias interbancarias, se implementaron actualizaciones tecnológicas para mejorar la precisión del sistema de asignación de cuentas, y se realizó capacitación y entrenamiento para los equipos responsables de gestionar estas transacciones.
- **Resultados Esperados:** Se espera que esta actividad mejore la eficiencia operativa en transferencias interbancarias, reduzca errores y retrasos en el procesamiento de las mismas, y fortalezca la confianza en el sistema de gestión de pagos.

**Título: Integración Cobre**
- **Contexto:** Este proyecto se enfocó en la incorporación de un nuevo servicio financiero, denominado "Cobre", que permitirá a los usuarios acceder a una variedad de servicios y funcionalidades bancarias desde una plataforma centralizada.
- **Plan de Rollout:** El plan de implementación para "Integración Cobre" incluyó etapas claras de desarrollo, pruebas y lanzamiento en diferentes segmentos del mercado objetivo. Se abordaron aspectos como la integración de APIs con sistemas existentes, la adaptación de interfaces de usuario, y la capacitación de personal para manejar las nuevas funcionalidades.
- **Objetivo:** El principal objetivo fue lanzar "Cobre" en un mercado competitivo, asegurando que los servicios ofrecidos cumplieran con estándares de calidad y seguridad exigidos por la institución financiera. Se buscaba no solo introducir un nuevo producto, sino también mejorar el posicionamiento estratégico de la entidad en el mercado mediante la mejora de la experiencia del cliente.
- **Progreso:** Durante el día anterior, se completaron las etapas iniciales de planificación y configuración técnica. Se realizó una evaluación detallada de los sistemas existentes para identificar puntos de integración y se iniciaron conversaciones con proveedores potenciales para asegurar un apoyo tecnológico robusto durante la fase de lanzamiento.
- **Resultados Esperados:** Los resultados esperados incluyen una mayor diversificación de servicios financieros, una mejora en la satisfacción del cliente y, finalmente, el aumento de la rentabilidad a través de un mejor aprovechamiento de los mercados y segmentos disponibles.
Resumen de las actividades - 2024-08-26 22:24:32
 Hace un día, se llevó a cabo una reunión en la que se discutieron y planearon varias actividades importantes para el futuro operativo de Cobre. En esta reunión, se acordó cambiar por defecto las transacciones realizadas en la plataforma Cobre hacia cuentas de ahorro, ya que es más común que los colaboradores utilicen este tipo de cuenta. Esta decisión tiene como objetivo facilitar y optimizar el proceso de manejo financiero para todos los usuarios.

Además, se acordó tomar medidas adicionales para restringir el acceso a las funcionalidades avanzadas como exportación, auditoría y conciliaciones. Esto se hizo con la intención de mejorar la seguridad y asegurar que solo aquellos con necesidades específicas puedan acceder a estas áreas avanzadas.

Estas decisiones reflejan el compromiso continuo de Cobre para brindar un servicio eficiente, seguro y accesible tanto para los usuarios como para los administradores del sistema.
Resumen de las actividades - 2024-08-26 22:27:32
 El resumen de las actividades del día anterior en Cobre se centra en dos aspectos principales: el cambio de la cuenta por defecto para todas las transacciones y la restricción del acceso a determinadas acciones como exportación, auditoría y conciliaciones.

**Cambio de Cuenta Por Defecto:**
En Cobre, se ha procedido al ajuste de que todas las transacciones por defecto sean realizadas desde cuentas de ahorro en lugar de cuentas corrientes. Esta decisión responde al uso predominante y la preferencia generalizada de los colaboradores por las cuentas de ahorro para manejar sus operaciones financieras. Este cambio facilitará el manejo de fondos y asegurará que los recursos estén disponibles en momentos oportunos, ya que las cuentas de ahorro suelen tener restricciones sobre la movilidad de los fondos.

**Restricción del Acceso a Acciones Específicas:**
Además de ajustar la cuenta por defecto, se ha implementado una medida para restringir el acceso a las funciones avanzadas como exportación, auditoría y conciliaciones. Esta restricción tiene como objetivo principal asegurar la integridad y seguridad de los datos financieros, evitando que usuarios no autorizados accedan a información sensible o realicen operaciones inapropiadas.

**Contexto del Segundo Aspecto:**
El segundo punto abarca un contexto detallado sobre las razones y motivaciones detrás de la restricción mencionada. Este aspecto incluye una explicación más profunda sobre por qué estas acciones específicas han sido restringidas, posibles riesgos o incidencias que pudieran presentarse sin esta medida de seguridad adicional, y cómo se planea supervisar y controlar este acceso restringido.

Estos dos elementos son fundamentales para garantizar un correcto funcionamiento operativo en Cobre, una gestión eficiente de los recursos financieros y un alto nivel de seguridad y control sobre la información sensible que maneja la organización.
Resumen de las actividades - 2024-08-26 22:28:23
 **Resumen de las Actividades del Día Anterior:**

**Cambiar a cuenta de ahorro por defecto en todas las transacciones de Cobre:**
- En la plataforma de Cobre, se ha implementado un cambio para que todas las transacciones sean realizadas por defecto en una cuenta de ahorro. Este ajuste responde al uso predominante de cuentas de ahorro entre los colaboradores. El propósito es optimizar y facilitar el manejo de fondos mediante la promoción del uso de esta herramienta financiera más segura y estratégica para todos los usuarios.

**Restringir Acceso a las Acciones de Exportación/Auditoría/Conciliaciones:**
- **Contexto:** Esta medida se ha tomado con el objetivo de mejorar la seguridad y controlar accesos no autorizados o innecesarios a información sensible relacionada con operaciones de exportación, auditorías internas y conciliaciones financieras. 
- **Alcance:** Se han restringido los permisos para acceder a estas secciones del sistema, asignando solo a aquellos roles administrativos que tienen una necesidad específica y documentada de realizar dichas acciones. Esto busca prevenir errores involuntales o intencionados en operaciones críticas y proteger la integridad de los datos financieros y operativos.

Estas actividades son parte del compromiso con la seguridad, la eficiencia y el control interno en las operaciones de Cobre, asegurando que tanto los usuarios como los datos estén protegidos contra riesgos innecesarios y contribuyendo a un entorno de trabajo más seguro y efectivo.
Resumen de las actividades - 2024-08-26 23:02:17
 El resumen de las actividades del día anterior en Cobre indica que, debido a la preferencia generalizada por parte de los colaboradores de usar una cuenta de ahorro para realizar transacciones, se ha decidido cambiar el método de pago predeterminado para todas las transacciones realizadas en el sistema. Esto facilitará y optimizará el proceso financiero tanto para los usuarios como para la administración del banco.

Además, se ha implementado una medida adicional para aumentar la seguridad y controlar el acceso a determinadas acciones sensibles dentro del área de exportación, auditoría y conciliaciones. Esta restricción tiene como objetivo prevenir errores no autorizados o intencionados que puedan afectar negativamente los procesos financieros y los datos confidenciales.

Estas decisiones reflejan un enfoque proactivo de Cobre para mejorar la eficiencia operativa, proteger la información sensible y satisfacer las necesidades específicas tanto de sus clientes como del equipo interno.