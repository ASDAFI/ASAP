/// <reference types="react" />
import { Line as D3Line } from 'd3-shape';
import { BumpCommonProps, BumpComputedSerie, BumpDatum, BumpSerieExtraProps } from './types';
interface LineProps<Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps> {
    serie: BumpComputedSerie<Datum, ExtraProps>;
    lineGenerator: D3Line<[number, number | null]>;
    yStep: number;
    isInteractive: BumpCommonProps<Datum, ExtraProps>['isInteractive'];
    onMouseEnter?: BumpCommonProps<Datum, ExtraProps>['onMouseEnter'];
    onMouseMove?: BumpCommonProps<Datum, ExtraProps>['onMouseMove'];
    onMouseLeave?: BumpCommonProps<Datum, ExtraProps>['onMouseLeave'];
    onClick?: BumpCommonProps<Datum, ExtraProps>['onClick'];
    setActiveSerieIds: (serieIds: string[]) => void;
    tooltip: BumpCommonProps<Datum, ExtraProps>['tooltip'];
}
export declare const Line: <Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps>({ serie, lineGenerator, yStep, isInteractive, onMouseEnter, onMouseMove, onMouseLeave, onClick, setActiveSerieIds, tooltip, }: LineProps<Datum, ExtraProps>) => JSX.Element;
export {};
//# sourceMappingURL=Line.d.ts.map