import datetime
import json
from xml.dom import minidom
from xml.etree import ElementTree
from xml.etree.ElementTree import Element, SubElement

# use requests from botocore so we don't need to package all of requests just for the dict
# noinspection PyPackageRequirements
from dateutil.relativedelta import relativedelta

from dict_util import CaseInsensitiveDict

# note voted out is not the same as new president is sworn in
it_can_leave = datetime.datetime(2021, 1, 20, 12, 0, 0, 0)


def get_time_delta():
    return relativedelta(it_can_leave, datetime.datetime.now())


def make_xml_response(pretty):
    diff = get_time_delta()
    root = Element('tillTrump')
    human = SubElement(root, 'humanReadable')
    human.text = make_meatbag_response()
    years = SubElement(root, 'years')
    years.text = str(diff.years)
    months = SubElement(root, 'months')
    months.text = str(diff.months)
    days = SubElement(root, 'days')
    days.text = str(diff.days)
    hours = SubElement(root, 'hours')
    hours.text = str(diff.hours)
    minutes = SubElement(root, 'minutes')
    minutes.text = str(diff.minutes)
    seconds = SubElement(root, 'seconds')
    seconds.text = str(diff.seconds)
    et = ElementTree.tostring(root)

    dom = minidom.parseString(et)
    pretty_xml_as_string = dom.toprettyxml()
    return pretty_xml_as_string if pretty else et.decode('utf-8')


def make_json_response(pretty):
    delta = get_time_delta()
    resp = {
        "humanReadable": make_meatbag_response(),
        'years': delta.years,
        'months': delta.months,
        'hours': delta.hours,
        'minutes': delta.minutes,
        'seconds': delta.seconds
    }
    args = {"obj": resp, "indent": 4, "sort_keys": True} if pretty else {'obj': resp}
    return json.dumps(**args)


def make_meatbag_response():
    delta = get_time_delta()
    return f"President Donald Trump will leave office January 20th, 2021 at noon." \
           f" This is only {delta.years} year, {delta.hours} hours, {delta.minutes}" \
           f" minutes and {delta.seconds} seconds away!"


def make_aws_response(response, status_code=200, content_type="plain/text"):
    return {
        "statusCode": status_code,
        "body": f"{response}\n",
        "headers": {
            "Content-Type": content_type
        }
    }


def handler(event, context):
    event = CaseInsensitiveDict(event)
    headers = CaseInsensitiveDict(event.get('headers') or {})
    accept = headers.get('accept') or ""
    event_path = CaseInsensitiveDict(event.get('requestContext', {})).get('path')
    pretty = event_path == '/pretty'
    if 'json' in accept.lower():
        resp = make_aws_response(make_json_response(pretty), content_type='application/json')
    elif 'xml' in accept.lower():
        resp = make_aws_response(make_xml_response(pretty), content_type="application/xml")
    else:
        resp = make_aws_response(make_meatbag_response())
    return resp


if __name__ == '__main__':
    print(get_time_delta())
    print(make_xml_response())
    print(make_json_response())
