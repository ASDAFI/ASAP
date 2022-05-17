/// <reference types="react" />
import { InheritedColorConfig } from '@nivo/colors';
import { AreaBumpComputedSerie, AreaBumpDatum, AreaBumpLabel, AreaBumpSerieExtraProps } from './types';
interface AreaLabelsProps<Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps> {
    getLabel: Exclude<AreaBumpLabel<Datum, ExtraProps>, false>;
    series: AreaBumpComputedSerie<Datum, ExtraProps>[];
    position: 'start' | 'end';
    padding: number;
    color: InheritedColorConfig<AreaBumpComputedSerie<Datum, ExtraProps>>;
}
export declare const AreasLabels: <Datum extends AreaBumpDatum, ExtraProps extends AreaBumpSerieExtraProps>({ getLabel, series, position, padding, color, }: AreaLabelsProps<Datum, ExtraProps>) => JSX.Element;
export {};
//# sourceMappingURL=AreasLabels.d.ts.map