/// <reference types="react" />
import { InheritedColorConfig } from '@nivo/colors';
import { AreaBumpAreaPoint, AreaBumpCommonProps, AreaBumpComputedSerie, AreaBumpDataProps, AreaBumpDatum, AreaBumpInterpolation, AreaBumpLabel, AreaBumpLabelData, AreaBumpSerieExtraProps, DefaultAreaBumpDatum } from './types';
export declare const useAreaBump: <Datum extends AreaBumpDatum = DefaultAreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps = Record<string, unknown>>({ data, width, height, align, spacing, xPadding, interpolation, colors, fillOpacity, activeFillOpacity, inactiveFillOpacity, borderWidth, activeBorderWidth, inactiveBorderWidth, borderColor, borderOpacity, activeBorderOpacity, inactiveBorderOpacity, isInteractive, defaultActiveSerieIds, }: {
    data: import("./types").AreaBumpSerie<Datum, ExtraProps>[];
    width: number;
    height: number;
    align: import("./types").AreaBumpAlign;
    spacing: number;
    xPadding: number;
    interpolation: AreaBumpInterpolation;
    colors: import("@nivo/colors").OrdinalColorScaleConfig<import("./types").AreaBumpSerie<Datum, ExtraProps>>;
    fillOpacity: number;
    activeFillOpacity: number;
    inactiveFillOpacity: number;
    borderWidth: number;
    activeBorderWidth: number;
    inactiveBorderWidth: number;
    borderColor: InheritedColorConfig<Omit<AreaBumpComputedSerie<Datum, ExtraProps>, "borderColor" | "borderWidth" | "fillOpacity" | "borderOpacity">>;
    borderOpacity: number;
    activeBorderOpacity: number;
    inactiveBorderOpacity: number;
    isInteractive: boolean;
    defaultActiveSerieIds: string[];
}) => {
    series: AreaBumpComputedSerie<Datum, ExtraProps>[];
    xScale: import("@nivo/scales").ScalePoint<Datum["x"]>;
    heightScale: import("@nivo/scales").ScaleLinear<number>;
    areaGenerator: import("d3-shape").Area<AreaBumpAreaPoint>;
    activeSerieIds: string[];
    setActiveSerieIds: import("react").Dispatch<import("react").SetStateAction<string[]>>;
};
export declare const useAreaBumpSerieHandlers: <Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps>({ serie, isInteractive, onMouseEnter, onMouseMove, onMouseLeave, onClick, setActiveSerieIds, tooltip, }: {
    serie: AreaBumpComputedSerie<Datum, ExtraProps>;
    isInteractive: boolean;
    onMouseEnter?: import("./types").AreaBumpMouseHandler<Datum, ExtraProps> | undefined;
    onMouseMove?: import("./types").AreaBumpMouseHandler<Datum, ExtraProps> | undefined;
    onMouseLeave?: import("./types").AreaBumpMouseHandler<Datum, ExtraProps> | undefined;
    onClick?: import("./types").AreaBumpMouseHandler<Datum, ExtraProps> | undefined;
    setActiveSerieIds: (serieIds: string[]) => void;
    tooltip: import("./types").AreaBumpAreaTooltip<Datum, ExtraProps>;
}) => {
    onMouseEnter: ((event: any) => void) | undefined;
    onMouseMove: ((event: any) => void) | undefined;
    onMouseLeave: ((event: any) => void) | undefined;
    onClick: ((event: any) => void) | undefined;
};
export declare const useAreaBumpSeriesLabels: <Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps>({ series, position, padding, color, getLabel, }: {
    series: AreaBumpComputedSerie<Datum, ExtraProps>[];
    position: 'start' | 'end';
    padding: number;
    color: InheritedColorConfig<AreaBumpComputedSerie<Datum, ExtraProps>>;
    getLabel: true | ((serie: import("./types").AreaBumpSerie<Datum, ExtraProps>) => string);
}) => AreaBumpLabelData<Datum, ExtraProps>[];
//# sourceMappingURL=hooks.d.ts.map