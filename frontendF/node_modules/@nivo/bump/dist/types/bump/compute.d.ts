import { BumpDataProps, BumpDatum, BumpComputedSerie, BumpSerieExtraProps } from './types';
export declare const computeSeries: <Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps>({ width, height, data, xPadding, xOuterPadding, yOuterPadding, }: {
    width: number;
    height: number;
    data: import("./types").BumpSerie<Datum, ExtraProps>[];
    xPadding: number;
    xOuterPadding: number;
    yOuterPadding: number;
}) => {
    series: Omit<BumpComputedSerie<Datum, ExtraProps>, "color" | "opacity" | "lineWidth">[];
    xScale: import("@nivo/scales").ScalePoint<Datum["x"]>;
    yScale: import("@nivo/scales").ScalePoint<number>;
};
//# sourceMappingURL=compute.d.ts.map