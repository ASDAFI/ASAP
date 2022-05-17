/// <reference types="react" />
import { BumpDatum, BumpPoint, BumpSerieExtraProps } from './types';
interface PointProps<Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps> {
    point: BumpPoint<Datum, ExtraProps>;
}
export declare const Point: <Datum extends BumpDatum, ExtraProps extends BumpSerieExtraProps>({ point, }: PointProps<Datum, ExtraProps>) => JSX.Element;
export {};
//# sourceMappingURL=Point.d.ts.map