# FHX

## Allgemeine Klasse
 
 
- Value
- Rectangle

``` go
// Es sind nicht alle Parameter in jeder Value vorhanden
type Value struct {
  Set string
  StringValue string
  Changeable bool
  Scalable bool
  CV string
  High int
  Low int
  Units string
  Description string
}
```
Koordinaten zum erstellen von einer Graphik
Nur über die Posiiton kann vieles zugeordnet werden
``` go
type Rectangle struct {
  X int
  Y int
  H int
  W int
}
```
``` go
type Position struct {
  X int
  Y int
}
```
## Grund Parameter
Die Grund Parameter einer FHX Datei

- Version
- Erstellt
- SCHEMA
- LOCALE
- Batch

```
/* Version: 10.3.1.3657.xr */
/* "05-Aug-2020 09:21:32" */

SCHEMA
 user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */
{
  VERSION=1253152407/* "17-Sep-2009 03:53:27" */
  MAJOR_VERSION=10
  MINOR_VERSION=3
  MAINTENANCE_VERSION=1
  BUILD_VERSION=3664
  BUILD_ID="xr"
  VERSION_STR="10.3.1.3664.xr"
  ONLINE_UPGRADE=F
}
LOCALE
 user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */
{
  LOCALE="English_United States.1252"
}
```

``` go
type FHX struct {
  Version string
  Create string
  Schema Schema
  Locale Locale
  Batch Batch
}
```

## SCHEMA
Aufbau vom Schema
```
SCHEMA
 user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */
{
  VERSION=1253152407/* "17-Sep-2009 03:53:27" */
  MAJOR_VERSION=10
  MINOR_VERSION=3
  MAINTENANCE_VERSION=1
  BUILD_VERSION=3664
  BUILD_ID="xr"
  VERSION_STR="10.3.1.3664.xr"
  ONLINE_UPGRADE=F
}
```
``` go
type Schema struct {
  User string
  Time time.Time
  SchemaDetail SchemaDetail  
}

type SchemaDetail struct {
  Version time.Time
  MajorVersion int
  MinorVersion int
  MaintenanceVersion int
  BuildVersion int
  BuildId string
  OnlineUpgrade bool
  VersionStr string
}
```

## Locale
```
LOCALE
 user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */
{
  LOCALE="English_United States.1252"
}
```

``` go
type Locale struct {
  User string
  Time time.Time
  LocaleDetail LocaleDetail  
}

type LocaleDetail struct {
  Locale string
}
```

## Batch

Hier kommt der eigentlich Block von dem Rezept. Der Aufbau und die Parameter von den Rezepten. Leider sind die orginale nicht darin enthalten. Diese werden als `ORIGIN=CONSTANT` angegeben. Diese muss man wiederum aus der OP auslesen. Dazu brauch ich alle Parameter der OP.

### Aufbau
- Batch
- Formula_Parameter
- Attribute_Instance
- PFC_Algoritm
```
BATCH_RECIPE NAME="RP_REIN_IODOS-Q2800-P105_WA" TYPE=PROCEDURE CATEGORY="Recipes/Procedures/REIN_IODOS-3"
 user="FLAMBRIGGE" time=1596611912/* "05-Aug-2020 09:18:32" */
{
  DESCRIPTION="Hier die Beschreibung 123"
  USE_EQUIPMENT_TRAINS=F
  AUTHOR="LSCHMID1"
  ABSTRACT=""
  BATCH_UNITS=""
  BATCH_LENGTH=""
  DEFAULT_BATCH_SIZE=100
  MINIMUM_BATCH_SIZE=1
  MAXIMUM_BATCH_SIZE=100
  PRODUCT_CODE=""
  PRODUCT_NAME=""
  RECIPE_APPROVAL_INFO=""
  VERSION="1"
  FORMULA_PARAMETER NAME="FP_ANZAHL_DEST" TYPE=BATCH_PARAMETER_INTEGER
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
  }
  ATTRIBUTE_INSTANCE NAME="FP_ANZAHL_DEST"
  {
    VALUE { DESCRIPTION="" HIGH=255 LOW=0 SCALABLE=F CV=3 UNITS="            " }
  }
  PFC_ALGORITHM
```
``` go
type Batch struct {
  Name string
  Time time.Time
  User string
  Type string // Procedure, Operation, Unitprocedure
  Category string // Ist der Pfad der Datei
  BatchDetail BatchDetail  
}

type BatchDetail struct {
  DESCRIPTION           string
  USE_EQUIPMENT_TRAIN   bool
  AUTHOR                string
  ABSTRACT              string
  BATCH_UNITS           string
  BATCH_LENGTH          string
  DEFAULT_BATCH_SIZE    int
  MINIMUM_BATCH_SIZE    int
  MAXIMUM_BATCH_SIZE    int
  PRODUCT_CODE          string
  PRODUCT_NAME          string
  RECIPE_APPROVAL_INFO  string
  VERSION               string
  FormPara              []FormulaParameter
  AttrInstance          []AttributeInstance
  PFC                   PFC
}
```

## Formula_Parameter

```
 FORMULA_PARAMETER NAME="FP_ANZAHL_DEST" TYPE=BATCH_PARAMETER_INTEGER
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
  }
```
``` go
  type FormulaParameter struct {
    Name            string
    Type            string // obs ein abgeleiteter Typ ist eine eine Eingabe
    FormParamDetail FormParamDetail
  }
   type FormParamDetail struct {
    Connection  string
    Rectangle   Rectangle
  }
```
## Attribute_Instance

```
ATTRIBUTE_INSTANCE NAME="FP_ANZAHL_DEST"
  {
    VALUE { DESCRIPTION="" HIGH=255 LOW=0 SCALABLE=F CV=3 UNITS="            " }
  }
```
``` go
    type AttributeInstance struct {
      Name    string
      AttrInstDetail AttrInstDetail
    }
    type AttrInstDetail struct {
      Value Value
    }
```

# PFC 

Der Hauptteil in dem die einzelnen Daten gespeichert werden

- GRAPHICS ALGORITHM
- STEP
- INITIAL_STEP
- TRANSITION
- STEP_TRANSITION_CONNECTION
- TRANSITION_STEP_CONNECTION
```
PFC_ALGORITHM
  {
    GRAPHICS ALGORITHM=PFC
    {
      TEXT_GRAPHIC
      {
        NAME="{714D9FFE-64CD-4984-8B4D-8D78B1A55003}"
        ORIGIN= { X=271 Y=434 }
        END= { X=356 Y=448 }
        TEXT="Destillieren P5500"
      }
    }
    STEP NAME="START" DEFINITION=""
    {
      DESCRIPTION=""
      RECTANGLE= { X=490 Y=10 H=40 W=100 }
      KEY_PARAMETER=""      
    }
    STEP NAME="UP_P105_1TR:1" DEFINITION="UP_P105_1TR"
    {
      DESCRIPTION=""
      RECTANGLE= { X=680 Y=1780 H=30 W=240 }
      ACQUIRE_UNIT=F
      RETAIN_UNIT=T
      STEP_PARAMETER NAME="FP_FSB_LI"
      {
        ORIGIN=CONSTANT
      }
      ATTRIBUTE_INSTANCE NAME="FP_TR_SENDER"
      {
        VALUE
        {
          SET="LP_UNITS"
          STRING_VALUE="Q2300"
          CHANGEABLE=F
        }
      }
    }
    INITIAL_STEP="START"
    TRANSITION NAME="T1"
    {
      POSITION= { X=510 Y=70 }
      TERMINATION=F
      EXPRESSION="TRUE"
    }
    STEP_TRANSITION_CONNECTION STEP="START" TRANSITION="T1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T1" STEP="UP_Q2800_START:1" { }
  }

``` 

``` go
    type PFC struct {
      GraphicsAlgorithm     GraphicsAlgorithm 
      Step                  []Step
      InitialStep           InitialStep
      Transition            []Transition
      StepTransition        []StepTransition
      TransitionStep        []TransitionStep 
    }
``` 

## GraphicsAlgorithm
Behinhaltet die Zeichnung Objekte und Graphischen Objekte
```
GRAPHICS ALGORITHM=PFC
    {
      TEXT_GRAPHIC
      {
        NAME="{714D9FFE-64CD-4984-8B4D-8D78B1A55003}"
        ORIGIN= { X=271 Y=434 }
        END= { X=356 Y=448 }
        TEXT="Destillieren P5500"
      }
    }
```
``` go
    type GraphicsAlgorithm struct {
      TextGraphic []TextGraphic
    }
    type TextGraphic struct {
      Name      string
      Origin    Position
      End       Position
      Text      string
    }
``` 
## Step
Step ist die UP oder OP mit den Parameter
```
STEP NAME="START" DEFINITION=""
    {
      DESCRIPTION=""
      RECTANGLE= { X=490 Y=10 H=40 W=100 }
      KEY_PARAMETER=""      
    }
    STEP NAME="UP_P105_1TR:1" DEFINITION="UP_P105_1TR"
    {
      DESCRIPTION=""
      RECTANGLE= { X=680 Y=1780 H=30 W=240 }
      ACQUIRE_UNIT=F
      RETAIN_UNIT=T
      STEP_PARAMETER NAME="FP_FSB_LI"
      {
        ORIGIN=CONSTANT
      }
      ATTRIBUTE_INSTANCE NAME="FP_TR_SENDER"
      {
        VALUE
        {
          SET="LP_UNITS"
          STRING_VALUE="Q2300"
          CHANGEABLE=F
        }
      }
    }
```

``` go
  type Step struct {
    Description         string
    Rectangle           Rectangle
    AquireUnit          bool
    RetainUnit          bool
    StepParameter       []StepParameter
    AttributeInstance   []AttributeInstance
  }  
  type StepParameter struct {
    Name    string
    Origin  string
    Deferred string
  } 
```

## InitialStep
Beinhaltet den Start Step
``` go
  type InitialStep struct {
    Step string
  }
```

## Transition
Verknüpfungen der einzelnen Steps und ihre Position
```
TRANSITION NAME="T1"
  {
    POSITION= { X=510 Y=70 }
    TERMINATION=F
    EXPRESSION="TRUE"
  }
```
``` go
  type Transition struct {
    Name string
    TransitionDetial TransitionDetail
  }
  type TransitionDetial struct {
    Position    Position
    Termination bool
    Expression  string
  }
```

## Transition Verknüpfungen
Einmal vom Step zur Transition und einmal von der Transition zum Step
```
    STEP_TRANSITION_CONNECTION STEP="START" TRANSITION="T1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T1" STEP="UP_Q2800_START:1" { }
```
``` go
  type TransitionStep struct {
    Name string
    TransitionDetial TransitionDetail
  }
  type StepTransition struct {
    Step        Step
    Transition  Transition
    Direction   bool // Vorwärts Rückwärts
  }  
```

## Vorgehen
### Einlesen fhx Datei
1. Kontrolle ob FHX
2. Ganze Datei in Zeilen aufsplitten
3. Datei in die Verschiedenen Kompontenten aufteilen.
  - 