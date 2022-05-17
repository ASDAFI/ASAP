/// <reference types="react" />
import { InheritedColorConfig } from '@nivo/colors';
import { BumpInterpolation, BumpCommonProps, BumpDatum, DefaultBumpDatum, BumpDataProps, BumpComputedSerie, BumpPoint, BumpLabel, BumpLabelData, BumpSerieExtraProps } from './types';
export declare const useBump: <Datum extends BumpDatum = DefaultBumpDatum, ExtraProps extends BumpSerieExtraProps = Record<string, never>>({ width, height, data, interpolation, xPadding, xOuterPadding, yOuterPadding, lineWidth, activeLineWidth, inactiveLineWidth, colors, opacity, activeOpacity, inactiveOpacity, pointSize, activePointSize, inactivePointSize, pointColor, pointBorderWidth, activePointBorderWidth, inactivePointBorderWidth, pointBorderColor, isInteractive, defaultActiveSerieIds, }: {
    width: number;
    height: number;
    data: import("./types").BumpSerie<Datum, ExtraProps>[];
    interpolation: BumpInterpolation;
    xPadding: number;
    xOuterPadding: number;
    yOuterPadding: number;
    lineWidth: number;
    activeLineWidth: number;
    inactiveLineWidth: number;
    colors: import("@nivo/colors").OrdinalColorScaleConfig<import("./types").BumpSerie<Datum, ExtraProps>>;
    opacity: number;
    activeOpacity: number;
    inactiveOpacity: number;
    pointSize: number;
    activePointSize: number;
    inactivePointSize: number;
    pointColor: InheritedColorConfig<Omit<BumpPoint<Datum, ExtraProps>, "color" | "borderColor">>;
    pointBorderWidth: number;
    activePointBorderWidth: number;
    inactivePointBorderWidth: number;
    pointBorderColor: InheritedColorConfig<Omit<BumpPoint<Datum, ExtraProps>, "borderColor">>;
    isInteractive: boolean;
    defaultActiveSerieIds: string[];
}) => {
    xScale: import("@nivo/scales").ScalePoint<Datum["x"]>;
    yScale: import("@nivo/scales").ScalePoint<number>;
    series: BumpComputedSerie<Datum, ExtraProps>[];
    points: BumpPoint<Datum, ExtraProps>[];
    lineGenerator: import("d3-shape").Line<[number, number | null]>;
    activeSerieIds: string[];
    setActiveSerieIds: import("react").Dispatch<import("react").SetStateAction<string[]>>;
};
export declare const useBumpSerieHandlers: <Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps>({ serie, isInteractive, onMouseEnter, onMouseMove, onMouseLeave, onClick, setActiveSerieIds, tooltip, }: {
    serie: BumpComputedSerie<Datum, ExtraProps>;
    isInteractive: boolean;
    onMouseEnter?: import("./types").BumpMouseHandler<Datum, ExtraProps> | undefined;
    onMouseMove?: import("./types").BumpMouseHandler<Datum, ExtraProps> | undefined;
    onMouseLeave?: import("./types").BumpMouseHandler<Datum, ExtraProps> | undefined;
    onClick?: import("./types").BumpMouseHandler<Datum, ExtraProps> | undefined;
    setActiveSerieIds: (serieIds: string[]) => void;
    tooltip: import("./types").BumpLineTooltip<Datum, ExtraProps>;
}) => {
    onMouseEnter: ((event: any) => void) | undefined;
    onMouseMove: ((event: any) => void) | undefined;
    onMouseLeave: ((event: any) => void) | undefined;
    onClick: ((event: any) => void) | undefined;
};
export declare const useBumpSeriesLabels: <Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps>({ series, position, padding, color, getLabel, }: {
    series: BumpComputedSerie<Datum, ExtraProps>[];
    position: 'start' | 'end';
    padding: number;
    color: InheritedColorConfig<BumpComputedSerie<Datum, ExtraProps>>;
    getLabel: true | ((serie: import("./types").BumpSerie<Datum, ExtraProps>) => string);
}) => BumpLabelData<Datum, ExtraProps>[];
//# sourceMappingURL=hooks.d.ts.map