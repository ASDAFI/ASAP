/// <reference types="react" />
import { AreaBumpDatum, AreaBumpComputedSerie, AreaBumpSerieExtraProps } from './types';
interface AreaTooltipProps<Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps> {
    serie: AreaBumpComputedSerie<Datum, ExtraProps>;
}
export declare const AreaTooltip: <Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps>({ serie, }: AreaTooltipProps<Datum, ExtraProps>) => JSX.Element;
export {};
//# sourceMappingURL=AreaTooltip.d.ts.map