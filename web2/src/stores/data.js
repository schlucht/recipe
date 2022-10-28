export const product = [
  {
    productName: "Asura",
    units: [
      {
        productID: "ASUR_Q2000",
        unitName: "Q2000",
        procedures: [
          {
            procedureName: "UP_Q2000_START:1",
            procedureDesc: "initialisiert die Unit",
            parameters: [
              {
                parameterName: "FP_APP_SOLE",
                paramterDesc: "Schaltet die Sole auf Warm",
                value: 45,
                unit: "°C",
                history: [
                  {
                    datum: "2022-2-12",
                    value: 35,
                    visum: "SL",
                  },
                ],
              },
            ],
          },
        ],
      },
    ],
  },
];
