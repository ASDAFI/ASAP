/// <reference types="react" />
import { InheritedColorConfig } from '@nivo/colors';
import { BumpComputedSerie, BumpDatum, BumpLabel, BumpSerieExtraProps } from './types';
interface LineLabelsProps<Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps> {
    series: BumpComputedSerie<Datum, ExtraProps>[];
    getLabel: Exclude<BumpLabel<Datum, ExtraProps>, false>;
    position: 'start' | 'end';
    padding: number;
    color: InheritedColorConfig<BumpComputedSerie<Datum, ExtraProps>>;
}
export declare const LinesLabels: <Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps>({ series, getLabel, position, padding, color, }: LineLabelsProps<Datum, ExtraProps>) => JSX.Element;
export {};
//# sourceMappingURL=LinesLabels.d.ts.map