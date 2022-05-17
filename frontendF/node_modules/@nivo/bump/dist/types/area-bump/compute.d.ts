import { ScalePoint, ScaleLinear } from '@nivo/scales';
import { AreaBumpCommonProps, AreaBumpComputedSerie, AreaBumpDataProps, AreaBumpDatum, AreaBumpSerieExtraProps } from './types';
export declare const computeSeries: <Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps>({ data, width, height, align, spacing, xPadding, }: {
    data: import("./types").AreaBumpSerie<Datum, ExtraProps>[];
    width: number;
    height: number;
    align: import("./types").AreaBumpAlign;
    spacing: number;
    xPadding: number;
}) => {
    series: Omit<AreaBumpComputedSerie<Datum, ExtraProps>, "color" | "borderColor" | "borderWidth" | "fill" | "fillOpacity" | "borderOpacity">[];
    xScale: ScalePoint<Datum["x"]>;
    heightScale: ScaleLinear<number>;
};
//# sourceMappingURL=compute.d.ts.map