/// <reference types="react" />
import { AreaBumpAreaGenerator, AreaBumpCommonProps, AreaBumpComputedSerie, AreaBumpDatum, AreaBumpSerieExtraProps } from './types';
interface AreaProps<Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps> {
    serie: AreaBumpComputedSerie<Datum, ExtraProps>;
    areaGenerator: AreaBumpAreaGenerator;
    blendMode: AreaBumpCommonProps<Datum, ExtraProps>['blendMode'];
    isInteractive: AreaBumpCommonProps<Datum, ExtraProps>['isInteractive'];
    onMouseEnter?: AreaBumpCommonProps<Datum, ExtraProps>['onMouseEnter'];
    onMouseMove?: AreaBumpCommonProps<Datum, ExtraProps>['onMouseMove'];
    onMouseLeave?: AreaBumpCommonProps<Datum, ExtraProps>['onMouseLeave'];
    onClick?: AreaBumpCommonProps<Datum, ExtraProps>['onClick'];
    setActiveSerieIds: (serieIds: string[]) => void;
    tooltip: AreaBumpCommonProps<Datum, ExtraProps>['tooltip'];
}
export declare const Area: <Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps>({ serie, areaGenerator, blendMode, isInteractive, onMouseEnter, onMouseMove, onMouseLeave, onClick, setActiveSerieIds, tooltip, }: AreaProps<Datum, ExtraProps>) => JSX.Element;
export {};
//# sourceMappingURL=Area.d.ts.map