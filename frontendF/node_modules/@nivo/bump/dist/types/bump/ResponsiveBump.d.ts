/// <reference types="react" />
import { BumpDatum, BumpSerieExtraProps, BumpSvgProps, DefaultBumpDatum } from './types';
export declare const ResponsiveBump: <Datum extends BumpDatum = DefaultBumpDatum, ExtraProps extends BumpSerieExtraProps = Record<string, never>>(props: Omit<BumpSvgProps<Datum, ExtraProps>, "width" | "height">) => JSX.Element;
//# sourceMappingURL=ResponsiveBump.d.ts.map