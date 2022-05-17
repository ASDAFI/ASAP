/// <reference types="react" />
import { BumpDatum, BumpPoint, BumpPointComponent, BumpSerieExtraProps } from './types';
interface PointsProps<Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps> {
    points: BumpPoint<Datum, ExtraProps>[];
    pointComponent: BumpPointComponent<Datum, ExtraProps>;
}
export declare const Points: <Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps>({ points, pointComponent, }: PointsProps<Datum, ExtraProps>) => JSX.Element;
export {};
//# sourceMappingURL=Points.d.ts.map