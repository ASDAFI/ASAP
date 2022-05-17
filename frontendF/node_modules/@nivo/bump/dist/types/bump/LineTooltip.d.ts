/// <reference types="react" />
import { BumpComputedSerie, BumpDatum, BumpSerieExtraProps } from './types';
interface LineTooltipProps<Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps> {
    serie: BumpComputedSerie<Datum, ExtraProps>;
}
export declare const LineTooltip: <Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps>({ serie, }: LineTooltipProps<Datum, ExtraProps>) => JSX.Element;
export {};
//# sourceMappingURL=LineTooltip.d.ts.map