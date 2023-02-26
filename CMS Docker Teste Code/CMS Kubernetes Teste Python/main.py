import time
import logging
from jaeger_client import Config # pip install jaeger-client
from jaeger_client.metrics.prometheus import PrometheusMetricsFactory # pip install prometheus_client
# from opentracing import Scope # pip install opentracing
# from opentracing.ext import tags as ext_tags
# pip install ipython
# pip install --upgrade "setuptools>=29" "pip>=9"


def init_jaeger_tracer(service_name: str):
	logging.getLogger('').handlers = []
	logging.basicConfig(format='[%(asctime)s,%(msecs)03d] | %(levelname)7s | %(module)10s |  %(lineno)4d | %(message)s', level=logging.DEBUG)
	config = Config(config={'sampler': {'type': 'const','param': 1,}, 'logging': True}, validate=True, service_name=service_name, metrics_factory=PrometheusMetricsFactory(service_name_label=service_name))
	return config.initialize_tracer()


if __name__ == "__main__":
	tracer = init_jaeger_tracer(service_name='cms-tnb-trace-1')
	try:

		logging.info(msg=f'INI')

		with tracer.start_span('MetodoMain') as span:
			span.set_tag('Movie1', 'ssssssssss')
			span.set_tag('Movie', 'ddd')
			span.log_kv({'event': 'CheckCinema' , 'value': 'dddddddddd'})
			span.set_tag('Movie2', 'ddd')
			span.log_kv({'event': 'CheckCinema' , 'value': 'dddddddddd'})
			span.log_kv({'event': 'CheckCinema' , 'value': 'dddddddddd'})
			span.log_kv({'event': 'CheckCinema' , 'value': 'dddddddddd'})
			span.log_kv({'event': 'CheckCinema' , 'value': 'dddddddddd'})
			span.log_kv({'eventdd': 'CheckCinemafff' , 'value': 'rrr'})

		# with tracer.start_span('CmsSpanPrinc') as span:
		# 	try:
		# 		span.set_tag('jaeger-debug-id', 'some-correlation-id')
		# 		# span.set_tag('CmsSpanChild1', 'some-correlation-id')
		# 		# span.set_tag('CmsSpanChildItem1', f"I'm working... #1")
		# 		# span.set_tag('CmsSpanChildItem2', f"I'm working... #2")
		# 		# span.set_tag('CmsSpanChildItem3', f"I'm working... #3")
		# 		span.log_kv({'event': 'test message', 'life': 42})
		# 		span.log_kv({'event': 'test message', 'life': 43})
		# 		with tracer.start_span('CmsSpanChild2', child_of=span) as child_span:
		# 			child_span.log_kv({'event': 'down below'})
		# 	except KeyboardInterrupt:
		# 		pass
		# 	except Exception as e:
		# 		span.set_tag('erro',str(e))

	finally:
		time.sleep(2)
		tracer.close()
		logging.info(msg=f'FIM')
